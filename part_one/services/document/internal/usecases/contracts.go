package usecases

import (
	"context"

	"github.com/valerii-smirnov/veryfi-test-task/part_one/document/internal/domain"

	"github.com/fsnotify/fsnotify"
	"github.com/veryfi/veryfi-go/veryfi/scheme"
)

type DocumentRepository interface {
	Create(ctx context.Context, document domain.Document) (domain.Document, error)
	Get(ctx context.Context, id uint) (domain.Document, error)
	Delete(ctx context.Context, id uint) error
	GetByFileName(ctx context.Context, fileName string) (domain.Document, error)
}

type VeryfiProcessor interface {
	ProcessDocumentUpload(opts scheme.DocumentUploadOptions) (*scheme.Document, error)
}

type FileSystemWatcher interface {
	GetEventsChan() <-chan fsnotify.Event
	GetErrorsChan() <-chan error
}

type EventProducer interface {
	ProduceCreate(ctx context.Context, id uint) error
	ProduceRemove(ctx context.Context, id uint) error
}
