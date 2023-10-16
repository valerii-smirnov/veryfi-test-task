package usecases

import (
	"context"
	"time"

	"github.com/valerii-smirnov/veryfi-test-task/part_one/stats/config"
	"github.com/valerii-smirnov/veryfi-test-task/part_one/stats/internal/domain"
	"github.com/valerii-smirnov/veryfi-test-task/part_one/stats/lib/event"

	"github.com/sirupsen/logrus"
	"github.com/veryfi/veryfi-go/veryfi/scheme"
)

type Processor struct {
	createdDocumentChan event.CreatedChan
	removedDocumentChan event.RemovedChan

	documentRepository    DocumentRepository
	geolocationRepository GeolocationRepository
	receiptRepository     ReceiptRepository
	config                config.Config
}

func NewProcessor(
	createdDocumentChan event.CreatedChan,
	removedDocumentChan event.RemovedChan,
	documentRepository DocumentRepository,
	geolocationRepository GeolocationRepository,
	receiptRepository ReceiptRepository,
	config config.Config,
) *Processor {
	return &Processor{
		createdDocumentChan:   createdDocumentChan,
		removedDocumentChan:   removedDocumentChan,
		documentRepository:    documentRepository,
		geolocationRepository: geolocationRepository,
		receiptRepository:     receiptRepository,
		config:                config,
	}
}

func (p Processor) Run(ctx context.Context) {
	for i := 0; i < p.config.EventProcessorWorkersPoolSize; i++ {
		go func(workerID int) {
			fields := logrus.Fields{
				"worker.id": workerID,
			}
			logrus.WithContext(ctx).WithFields(fields).Info("starting event processing worker...")
			defer logrus.WithContext(ctx).WithFields(fields).Info("stopping event processing worker...")

			for {
				select {
				case id := <-p.createdDocumentChan:
					logrus.WithContext(ctx).
						WithFields(fields).WithField("document.id", id).
						Info("received new created document event")

					if err := p.processCreatedDocumentEvent(ctx, id); err != nil {
						logrus.WithContext(ctx).WithError(err).Error("document created event processing error")
					}
				case id := <-p.removedDocumentChan:
					logrus.WithContext(ctx).
						WithFields(fields).WithField("document.id", id).
						Info("received new removed document event")

					if err := p.processRemovedDocumentEvent(ctx, id); err != nil {
						logrus.WithContext(ctx).WithError(err).Error("document removed event processing error")
					}
				case <-ctx.Done():
					return
				}
			}
		}(i)
	}
}

func (p Processor) processCreatedDocumentEvent(ctx context.Context, documentID uint) error {
	document, err := p.documentRepository.Get(ctx, documentID)
	if err != nil {
		return err
	}

	receipt, err := p.buildReceipt(ctx, documentID, document.VeryfiDocumentInfo)
	if err != nil {
		return err
	}

	if _, err = p.receiptRepository.Save(ctx, receipt); err != nil {
		return err
	}

	return nil
}

func (p Processor) processRemovedDocumentEvent(ctx context.Context, documentID uint) error {
	if err := p.receiptRepository.DeleteByDocumentID(ctx, documentID); err != nil {
		return err
	}

	return nil
}

func (p Processor) buildReceipt(ctx context.Context, documentID uint, document *scheme.Document) (domain.Receipt, error) {
	receiptDate, err := time.Parse(time.DateTime, document.Date)
	if err != nil {
		return domain.Receipt{}, err
	}

	items := make([]domain.Item, 0, len(document.LineItems))
	for _, lineItem := range document.LineItems {
		items = append(items, domain.Item{
			Type:        lineItem.Type,
			Description: lineItem.Description,
			Discount:    lineItem.Discount,
			Quantity:    lineItem.Quantity,
			Tax:         lineItem.Tax,
			TaxRate:     lineItem.TaxRate,
			Total:       lineItem.Total,
		})
	}

	var geography *domain.Geography

	if document.Vendor.Address != "" {
		geography, err = p.geolocationRepository.GetLocation(ctx, document.Vendor.Address)
		if err != nil {
			return domain.Receipt{}, err
		}
	}

	receipt := domain.Receipt{
		DocumentID:    documentID,
		Category:      document.Category,
		InvoiceNumber: document.InvoiceNumber,
		Currency:      document.CurrencyCode,
		Tax:           document.Tax,
		Total:         document.Total,
		Date:          receiptDate,
		Geography:     geography,
		Items:         items,
	}

	return receipt, nil
}
