package usecases

import (
	"context"

	"github.com/valerii-smirnov/veryfi-test-task/part_one/stats/internal/domain"
)

//go:generate mockgen -destination=./mock_test.go -package=usecases -source=./contracts.go

type DocumentRepository interface {
	Get(ctx context.Context, id uint) (domain.Document, error)
}

type GeolocationRepository interface {
	GetLocation(ctx context.Context, address string) (*domain.Geography, error)
}

type ReceiptRepository interface {
	Save(ctx context.Context, receipt domain.Receipt) (domain.Receipt, error)
	GetByDocumentID(ctx context.Context, documentID uint) (domain.Receipt, error)
	DeleteByDocumentID(ctx context.Context, documentID uint) error
}

type StatsRepository interface {
	GetTotalTaxByPeriod(ctx context.Context, period domain.Period) (domain.TotalTax, error)
	GetTotalDiscountByPeriod(ctx context.Context, period domain.Period) (domain.TotalDiscount, error)
	GetGeographyByPeriod(ctx context.Context, period domain.Period) ([]domain.GeographyInfo, error)
}
