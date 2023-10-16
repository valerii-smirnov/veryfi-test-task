package repositories

import (
	"context"

	"github.com/valerii-smirnov/veryfi-test-task/part_one/stats/internal/domain"
	"github.com/valerii-smirnov/veryfi-test-task/part_one/stats/internal/repositories/models"

	"gorm.io/gorm"
)

type Stats struct {
	db *gorm.DB
}

func NewStats(db *gorm.DB) *Stats {
	return &Stats{
		db: db,
	}
}

func (s Stats) GetTotalTaxByPeriod(ctx context.Context, period domain.Period) (domain.TotalTax, error) {
	var totalTax float64

	err := s.db.WithContext(ctx).
		Table(models.ReceiptTableName).
		Select("SUM(tax) as totalTax").
		Where("date BETWEEN ? AND ?", period.From, period.To).
		Where("deleted_at IS NULL").
		Scan(&totalTax).Error

	if err != nil {
		return domain.TotalTax{}, err
	}

	return domain.TotalTax{
		Total: totalTax,
	}, nil
}

func (s Stats) GetTotalDiscountByPeriod(ctx context.Context, period domain.Period) (domain.TotalDiscount, error) {
	var totalDiscount float64

	err := s.db.WithContext(ctx).
		Table(models.ItemTableName+" i").
		Select("SUM(i.discount) as totalDiscount").
		Joins("INNER JOIN receipt r ON i.receipt_id = r.id").
		Where("i.deleted_at IS NULL AND r.deleted_at IS NULL").
		Where("r.date BETWEEN ? AND ?", period.From, period.To).
		Scan(&totalDiscount).
		Error

	if err != nil {
		return domain.TotalDiscount{}, err
	}

	return domain.TotalDiscount{
		Total: totalDiscount,
	}, nil
}

func (s Stats) GetGeographyByPeriod(ctx context.Context, period domain.Period) ([]domain.GeographyInfo, error) {
	var receipts []models.Receipt
	err := s.db.WithContext(ctx).
		Preload("Geography", "deleted_at IS NULL").
		Where("date BETWEEN ? AND ?", period.From, period.To).
		Where("deleted_at IS NULL").
		Find(&receipts).Error

	if err != nil {
		return nil, err
	}

	domainGeos := make([]domain.GeographyInfo, 0, len(receipts))
	for _, receipt := range receipts {
		domainGeos = append(domainGeos, domain.GeographyInfo{
			Address:    receipt.Geography.Address,
			Lat:        receipt.Geography.Lat,
			Lng:        receipt.Geography.Lng,
			TotalPrice: receipt.Total,
		})
	}

	return domainGeos, nil
}
