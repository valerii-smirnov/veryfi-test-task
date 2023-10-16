package repositories

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	mgorm "github.com/valerii-smirnov/veryfi-test-task/part_one/pkg/testing/gorm"
	"github.com/valerii-smirnov/veryfi-test-task/part_one/stats/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
	"time"
)

func TestStats_GetTotalTaxByPeriod(t *testing.T) {
	db, mock, err := mgorm.NewMockedGorm(mgorm.WithLogLevel(logger.Info))
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.TODO()
	testingError := errors.New("testing error")
	from, _ := time.Parse(time.DateTime, "2023-10-10 00:00:00")
	to, _ := time.Parse(time.DateTime, "2023-10-14 00:00:00")

	period := domain.Period{
		From: from,
		To:   to,
	}

	tax := 333.33

	successTotalTax := domain.TotalTax{
		Total: tax,
	}

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx    context.Context
		period domain.Period
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		initMocks func()
		want      domain.TotalTax
		wantErr   bool
	}{
		{
			name: "db error",
			fields: fields{
				db: db,
			},
			args: args{
				ctx:    ctx,
				period: period,
			},
			initMocks: func() {
				mock.ExpectQuery("SELECT (.+) FROM `receipt`").WillReturnError(testingError)
			},
			want:    domain.TotalTax{},
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				db: db,
			},
			args: args{
				ctx:    ctx,
				period: period,
			},
			initMocks: func() {
				mock.ExpectQuery("SELECT (.+) FROM `receipt`").WillReturnRows(
					sqlmock.NewRows([]string{"totalTax"}).
						AddRow(tax),
				)
			},
			want:    successTotalTax,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.initMocks()
			s := NewStats(tt.fields.db)
			got, err := s.GetTotalTaxByPeriod(tt.args.ctx, tt.args.period)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestStats_GetTotalDiscountByPeriod(t *testing.T) {
	db, mock, err := mgorm.NewMockedGorm(mgorm.WithLogLevel(logger.Info))
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.TODO()
	testingError := errors.New("testing error")
	from, _ := time.Parse(time.DateTime, "2023-10-10 00:00:00")
	to, _ := time.Parse(time.DateTime, "2023-10-14 00:00:00")

	period := domain.Period{
		From: from,
		To:   to,
	}

	discount := 333.33

	successTotalTax := domain.TotalDiscount{
		Total: discount,
	}

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx    context.Context
		period domain.Period
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		initMocks func()
		want      domain.TotalDiscount
		wantErr   bool
	}{
		{
			name: "db error",
			fields: fields{
				db: db,
			},
			args: args{
				ctx:    ctx,
				period: period,
			},
			initMocks: func() {
				mock.ExpectQuery("SELECT (.+) FROM receipt_item i").WillReturnError(testingError)
			},
			want:    domain.TotalDiscount{},
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				db: db,
			},
			args: args{
				ctx:    ctx,
				period: period,
			},
			initMocks: func() {
				mock.ExpectQuery("SELECT (.+) FROM receipt_item i").WillReturnRows(
					sqlmock.NewRows([]string{"totalDiscount"}).AddRow(discount),
				)
			},
			want:    successTotalTax,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.initMocks()
			s := NewStats(tt.fields.db)
			got, err := s.GetTotalDiscountByPeriod(tt.args.ctx, tt.args.period)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestStats_GetGeographyByPeriod(t *testing.T) {
	db, mock, err := mgorm.NewMockedGorm(mgorm.WithLogLevel(logger.Info))
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.TODO()
	testingError := errors.New("testing error")
	from, _ := time.Parse(time.DateTime, "2023-10-10 00:00:00")
	to, _ := time.Parse(time.DateTime, "2023-10-14 00:00:00")

	period := domain.Period{
		From: from,
		To:   to,
	}

	expectedGeoInfo := []domain.GeographyInfo{
		{
			Address:    "Address 1",
			Lat:        12.2,
			Lng:        12.2,
			TotalPrice: 50,
		},
		{
			Address:    "Address 2",
			Lat:        44.4,
			Lng:        55.5,
			TotalPrice: 100,
		},
	}

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx    context.Context
		period domain.Period
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		initMocks func()
		want      []domain.GeographyInfo
		wantErr   bool
	}{
		{
			name: "query error",
			fields: fields{
				db: db,
			},
			args: args{
				ctx:    ctx,
				period: period,
			},
			initMocks: func() {
				mock.ExpectQuery("SELECT (.+) FROM `receipt`").WillReturnError(testingError)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				db: db,
			},
			args: args{
				ctx:    ctx,
				period: period,
			},
			initMocks: func() {
				mock.ExpectQuery("SELECT (.+) FROM `receipt`").WillReturnRows(
					sqlmock.NewRows([]string{"id", "total"}).
						AddRow(1, expectedGeoInfo[0].TotalPrice).
						AddRow(2, expectedGeoInfo[1].TotalPrice),
				)

				mock.ExpectQuery("SELECT (.+) FROM `receipt_geography`").WillReturnRows(
					sqlmock.NewRows([]string{"id", "receipt_id", "address", "lat", "lng"}).
						AddRow(1, 1, expectedGeoInfo[0].Address, expectedGeoInfo[0].Lat, expectedGeoInfo[0].Lng).
						AddRow(2, 2, expectedGeoInfo[1].Address, expectedGeoInfo[1].Lat, expectedGeoInfo[1].Lng),
				)
			},
			want:    expectedGeoInfo,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.initMocks()
			s := NewStats(tt.fields.db)
			got, err := s.GetGeographyByPeriod(tt.args.ctx, tt.args.period)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}
