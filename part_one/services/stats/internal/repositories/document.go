package repositories

import (
	"context"
	"encoding/json"

	"github.com/valerii-smirnov/veryfi-test-task/part_one/pkg/grpc/stubs/document"
	"github.com/valerii-smirnov/veryfi-test-task/part_one/stats/internal/domain"

	"github.com/veryfi/veryfi-go/veryfi/scheme"
)

func NewDocument(client document.DocumentServiceClient) *Document {
	return &Document{
		client: client,
	}
}

type Document struct {
	client document.DocumentServiceClient
}

func (d Document) Get(ctx context.Context, id uint) (domain.Document, error) {
	resp, err := d.client.GetDocumentByID(ctx, &document.GetDocumentByIDRequest{Id: uint64(id)})
	if err != nil {
		return domain.Document{}, err
	}

	var info scheme.Document
	if err := json.Unmarshal([]byte(resp.Document.VeryfiDocumentInfo), &info); err != nil {
		return domain.Document{}, err
	}

	doc := domain.Document{
		ID:                 uint(resp.Document.Id),
		VeryfiDocumentID:   uint(resp.Document.VeryfiDocumentID),
		VeryfiDocumentInfo: &info,
	}

	return doc, err
}
