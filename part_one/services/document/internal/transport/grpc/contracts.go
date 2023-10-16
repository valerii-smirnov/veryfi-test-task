package grpc

import (
	"context"

	"github.com/valerii-smirnov/veryfi-test-task/part_one/document/internal/domain"
)

type DocumentUseCase interface {
	GetDocumentByID(ctx context.Context, id uint) (domain.Document, error)
}
