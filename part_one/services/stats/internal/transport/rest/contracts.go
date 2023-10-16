package rest

import (
	"context"
	"time"

	"github.com/valerii-smirnov/veryfi-test-task/part_one/stats/internal/domain"
)

//go:generate mockgen -destination=./mock_test.go -package=rest -source=./contracts.go

type StatsUseCase interface {
	TotalTaxByPeriod(ctx context.Context, from, to time.Time) (domain.TotalTax, error)
	TotalDiscountByPeriod(ctx context.Context, from, to time.Time) (domain.TotalDiscount, error)
	GetGeographyByPeriod(ctx context.Context, from, to time.Time) ([]domain.GeographyInfo, error)
}
