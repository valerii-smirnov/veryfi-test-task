package usecases

import (
	"context"
	"github.com/valerii-smirnov/veryfi-test-task/part_one/document/internal/domain"
)

type Document struct {
	documentRepository DocumentRepository
}

func NewDocument(documentRepository DocumentRepository) *Document {
	return &Document{
		documentRepository: documentRepository,
	}
}

func (d Document) GetDocumentByID(ctx context.Context, id uint) (domain.Document, error) {
	return d.documentRepository.Get(ctx, id)
}
