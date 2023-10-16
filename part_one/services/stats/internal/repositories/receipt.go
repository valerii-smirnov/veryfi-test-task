package repositories

import (
	"context"

	"github.com/valerii-smirnov/veryfi-test-task/part_one/stats/internal/domain"
	"github.com/valerii-smirnov/veryfi-test-task/part_one/stats/internal/repositories/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Receipt struct {
	db *gorm.DB
}

func NewReceipt(db *gorm.DB) *Receipt {
	return &Receipt{
		db: db,
	}
}

func (r Receipt) Save(ctx context.Context, receipt domain.Receipt) (domain.Receipt, error) {
	items := make([]models.Item, 0, len(receipt.Items))
	for _, domainItem := range receipt.Items {
		items = append(items, models.Item{
			Type:        domainItem.Type,
			Description: domainItem.Description,
			Discount:    domainItem.Discount,
			Quantity:    domainItem.Quantity,
			Tax:         domainItem.Tax,
			TaxRate:     domainItem.TaxRate,
			Total:       domainItem.Total,
		})
	}

	receiptModel := models.Receipt{
		DocumentID:    receipt.DocumentID,
		Category:      receipt.Category,
		InvoiceNumber: receipt.InvoiceNumber,
		Currency:      receipt.Currency,
		Tax:           receipt.Tax,
		Total:         receipt.Total,
		Date:          receipt.Date,
		Geography: &models.Geography{
			Address: receipt.Geography.Address,
			Lat:     receipt.Geography.Lat,
			Lng:     receipt.Geography.Lng,
		},
		Items: items,
	}

	if err := r.db.WithContext(ctx).Save(&receiptModel).Error; err != nil {
		return domain.Receipt{}, err
	}

	return receipt, nil
}

func (r Receipt) GetByDocumentID(ctx context.Context, documentID uint) (domain.Receipt, error) {
	var receiptModel models.Receipt
	if err := r.db.WithContext(ctx).
		Preload("Geography", "deleted_at IS NOT NULL").
		Preload("Items", "deleted_at IS NOT NULL").
		First(&receiptModel, "document_id = ?", documentID).
		Error; err != nil {
		return domain.Receipt{}, err
	}

	domainItems := make([]domain.Item, 0, len(receiptModel.Items))
	for _, modelItem := range receiptModel.Items {
		domainItems = append(domainItems, domain.Item{
			ID:          modelItem.ID,
			Type:        modelItem.Type,
			Description: modelItem.Description,
			Discount:    modelItem.Discount,
			Quantity:    modelItem.Quantity,
			Tax:         modelItem.Tax,
			TaxRate:     modelItem.TaxRate,
			Total:       modelItem.Total,
			CreatedAt:   modelItem.CreatedAt,
			UpdatedAt:   modelItem.UpdatedAt,
		})
	}

	domainReceipt := domain.Receipt{
		ID:            receiptModel.ID,
		DocumentID:    receiptModel.DocumentID,
		Category:      receiptModel.Category,
		InvoiceNumber: receiptModel.InvoiceNumber,
		Currency:      receiptModel.Currency,
		Tax:           receiptModel.Tax,
		Total:         receiptModel.Total,
		Date:          receiptModel.Date,
		CreatedAt:     receiptModel.CreatedAt,
		UpdatedAt:     receiptModel.UpdatedAt,
		Geography: &domain.Geography{
			ID:        receiptModel.Geography.ID,
			Address:   receiptModel.Geography.Address,
			Lat:       receiptModel.Geography.Lat,
			Lng:       receiptModel.Geography.Lng,
			CreatedAt: receiptModel.Geography.CreatedAt,
			UpdatedAt: receiptModel.Geography.UpdatedAt,
		},
		Items: domainItems,
	}

	return domainReceipt, nil
}

func (r Receipt) DeleteByDocumentID(ctx context.Context, documentID uint) error {
	var receiptModel models.Receipt
	if err := r.db.WithContext(ctx).
		Preload("Geography", "deleted_at IS NULL").
		Preload("Items", "deleted_at IS NULL").
		First(&receiptModel, "document_id = ?", documentID).
		Error; err != nil {
		return err
	}

	if err := r.db.Select(clause.Associations).Delete(&receiptModel).Error; err != nil {
		return err
	}

	return nil
}
