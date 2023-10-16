package usecases

import (
	"context"
	"github.com/valerii-smirnov/veryfi-test-task/part_one/document/config"
	"github.com/valerii-smirnov/veryfi-test-task/part_one/document/internal/domain"
	"github.com/veryfi/veryfi-go/veryfi/scheme"
	"os"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
)

type Processor struct {
	watcher         FileSystemWatcher
	veryfiProcessor VeryfiProcessor
	eventProducer   EventProducer
	config          config.Config

	documentRepository DocumentRepository
}

func NewProcessor(
	watcher FileSystemWatcher,
	veryfiProcessor VeryfiProcessor,
	eventProducer EventProducer,
	cfg config.Config,
	fileRepository DocumentRepository,
) *Processor {
	return &Processor{
		watcher:            watcher,
		veryfiProcessor:    veryfiProcessor,
		eventProducer:      eventProducer,
		config:             cfg,
		documentRepository: fileRepository,
	}
}

func (p Processor) Run(ctx context.Context) error {
	for i := 0; i < p.config.WatcherWorkerPoolSize; i++ {
		go func(worker int) {

			logrus.Infof("running file processor worker #%d...", worker)
			defer logrus.Infof("stopping file processor worker #%d...", worker)
			for {
				select {
				case event := <-p.watcher.GetEventsChan():
					if err := p.processEvent(ctx, event, worker); err != nil {
						logrus.WithError(err).
							Error("file watcher event processing error")
					}
				case err := <-p.watcher.GetErrorsChan():
					logrus.WithError(err).
						Error("watching filesystem error occurred")
				case <-ctx.Done():
					return
				}
			}
		}(i)
	}

	return nil
}

func (p Processor) processEvent(ctx context.Context, event fsnotify.Event, worker int) error {
	switch event.Op {
	case fsnotify.Create:
		return p.processCreateEvent(ctx, event, worker)
	case fsnotify.Remove:
		return p.processRemoveDocument(ctx, event, worker)
	}

	return nil
}

func (p Processor) processCreateEvent(ctx context.Context, event fsnotify.Event, worker int) error {
	info, err := os.Stat(event.Name)
	if err != nil {
		return err
	}

	if info.IsDir() {
		return nil
	}

	fields := logrus.Fields{
		"worker":     worker,
		"event_name": event.Name,
	}

	logrus.WithFields(fields).Info("processing new document...")

	doc, err := p.veryfiProcessor.ProcessDocumentUpload(scheme.DocumentUploadOptions{
		FilePath: event.Name,
	})

	if err != nil {
		logrus.WithFields(fields).WithError(err).Error("new document processing error")
		return err
	}

	document := domain.Document{
		FileName:           strings.TrimPrefix(event.Name, p.config.FolderPath+"/"),
		VeryfiDocumentID:   uint(doc.ID),
		VeryfiDocumentInfo: doc,
	}

	document, err = p.documentRepository.Create(context.Background(), document)
	if err != nil {
		logrus.WithFields(fields).
			WithError(err).
			Error("storing document processing result into DB error")

		return err
	}

	if err := p.eventProducer.ProduceCreate(ctx, document.ID); err != nil {
		logrus.WithFields(fields).
			WithError(err).
			Error("event producing error")

		return err
	}

	logrus.WithFields(fields).Info("new document successfully processed")

	return nil
}

func (p Processor) processRemoveDocument(ctx context.Context, event fsnotify.Event, worker int) error {
	fields := logrus.Fields{
		"worker":     worker,
		"event_name": event.Name,
	}

	logrus.WithFields(fields).Info("processing removed document...")

	fileName := strings.TrimPrefix(event.Name, p.config.FolderPath+"/")
	document, err := p.documentRepository.GetByFileName(context.Background(), fileName)
	if err != nil {
		return err
	}

	if err := p.documentRepository.Delete(context.Background(), document.ID); err != nil {
		return err
	}

	if err := p.eventProducer.ProduceRemove(ctx, document.ID); err != nil {
		return err
	}

	logrus.WithFields(fields).Info("removed document successfully processed")

	return nil
}
