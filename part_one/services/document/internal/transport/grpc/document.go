package grpc

import (
	"context"
	"encoding/json"

	"github.com/valerii-smirnov/veryfi-test-task/part_one/pkg/grpc/stubs/document"
)

type Document struct {
	documentUseCase DocumentUseCase

	document.UnimplementedDocumentServiceServer
}

func NewDocument(documentUseCase DocumentUseCase) *Document {
	return &Document{
		documentUseCase: documentUseCase,
	}
}

func (d Document) GetDocumentByID(ctx context.Context, request *document.GetDocumentByIDRequest) (*document.GetDocumentByIDResponse, error) {
	doc, err := d.documentUseCase.GetDocumentByID(ctx, uint(request.Id))
	if err != nil {
		return nil, err
	}

	info, err := json.Marshal(doc.VeryfiDocumentInfo)
	if err != nil {
		return nil, err
	}

	return &document.GetDocumentByIDResponse{
		Document: &document.Document{
			Id:                 uint64(doc.ID),
			VeryfiDocumentID:   uint64(doc.VeryfiDocumentID),
			VeryfiDocumentInfo: string(info),
		},
	}, nil
}
