package usecases

import (
	"context"
	"time"

	"github.com/valerii-smirnov/veryfi-test-task/part_one/stats/internal/domain"
)

type Stats struct {
	statsRepository StatsRepository
}

func NewStats(statsRepository StatsRepository) *Stats {
	return &Stats{
		statsRepository: statsRepository,
	}
}

func (s Stats) TotalTaxByPeriod(ctx context.Context, from, to time.Time) (domain.TotalTax, error) {
	totalTax, err := s.statsRepository.GetTotalTaxByPeriod(ctx, domain.Period{
		From: from,
		To:   to,
	})

	if err != nil {
		return domain.TotalTax{}, err
	}

	return totalTax, nil
}

func (s Stats) TotalDiscountByPeriod(ctx context.Context, from, to time.Time) (domain.TotalDiscount, error) {
	totalDiscount, err := s.statsRepository.GetTotalDiscountByPeriod(ctx, domain.Period{
		From: from,
		To:   to,
	})

	if err != nil {
		return domain.TotalDiscount{}, err
	}

	return totalDiscount, nil
}

func (s Stats) GetGeographyByPeriod(ctx context.Context, from, to time.Time) ([]domain.GeographyInfo, error) {
	geos, err := s.statsRepository.GetGeographyByPeriod(ctx, domain.Period{
		From: from,
		To:   to,
	})

	if err != nil {
		return nil, err
	}

	return geos, nil
}
