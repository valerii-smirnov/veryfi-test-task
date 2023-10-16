package usecases

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/valerii-smirnov/veryfi-test-task/part_one/stats/internal/domain"
	"testing"
	"time"
)

func TestStats_TotalTaxByPeriod(t *testing.T) {
	controller := gomock.NewController(t)
	statsRepositoryMock := NewMockStatsRepository(controller)

	ctx := context.TODO()
	testingError := errors.New("testing error")
	from, _ := time.Parse(time.DateTime, "2023-10-10 00:00:00")
	to, _ := time.Parse(time.DateTime, "2023-10-14 00:00:00")
	period := domain.Period{
		From: from,
		To:   to,
	}

	successTotalTax := domain.TotalTax{
		Total: 5.55,
	}

	type fields struct {
		statsRepository StatsRepository
	}
	type args struct {
		ctx  context.Context
		from time.Time
		to   time.Time
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
			name: "repository error",
			fields: fields{
				statsRepository: statsRepositoryMock,
			},
			args: args{
				ctx:  ctx,
				from: from,
				to:   to,
			},
			initMocks: func() {
				statsRepositoryMock.EXPECT().GetTotalTaxByPeriod(gomock.Any(), gomock.Eq(period)).Return(domain.TotalTax{}, testingError)
			},
			want:    domain.TotalTax{},
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				statsRepository: statsRepositoryMock,
			},
			args: args{
				ctx:  ctx,
				from: from,
				to:   to,
			},
			initMocks: func() {
				statsRepositoryMock.EXPECT().GetTotalTaxByPeriod(gomock.Any(), gomock.Eq(period)).Return(successTotalTax, nil)
			},
			want:    successTotalTax,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.initMocks()
			s := NewStats(tt.fields.statsRepository)
			got, err := s.TotalTaxByPeriod(tt.args.ctx, tt.args.from, tt.args.to)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestStats_TotalDiscountByPeriod(t *testing.T) {
	controller := gomock.NewController(t)
	statsRepositoryMock := NewMockStatsRepository(controller)

	ctx := context.TODO()
	testingError := errors.New("testing error")
	from, _ := time.Parse(time.DateTime, "2023-10-10 00:00:00")
	to, _ := time.Parse(time.DateTime, "2023-10-14 00:00:00")

	period := domain.Period{
		From: from,
		To:   to,
	}

	successTotalDiscount := domain.TotalDiscount{
		Total: 100.22,
	}

	type fields struct {
		statsRepository StatsRepository
	}
	type args struct {
		ctx  context.Context
		from time.Time
		to   time.Time
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
			name: "repository error",
			fields: fields{
				statsRepository: statsRepositoryMock,
			},
			args: args{
				ctx:  ctx,
				from: from,
				to:   to,
			},
			initMocks: func() {
				statsRepositoryMock.EXPECT().GetTotalDiscountByPeriod(gomock.Any(), gomock.Eq(period)).Return(domain.TotalDiscount{}, testingError)
			},
			want:    domain.TotalDiscount{},
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				statsRepository: statsRepositoryMock,
			},
			args: args{
				ctx:  ctx,
				from: from,
				to:   to,
			},
			initMocks: func() {
				statsRepositoryMock.EXPECT().GetTotalDiscountByPeriod(gomock.Any(), gomock.Eq(period)).Return(successTotalDiscount, nil)
			},
			want:    successTotalDiscount,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.initMocks()
			s := NewStats(tt.fields.statsRepository)
			got, err := s.TotalDiscountByPeriod(tt.args.ctx, tt.args.from, tt.args.to)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestStats_GetGeographyByPeriod(t *testing.T) {
	controller := gomock.NewController(t)
	statsRepositoryMock := NewMockStatsRepository(controller)

	ctx := context.TODO()
	testingError := errors.New("testing error")
	from, _ := time.Parse(time.DateTime, "2023-10-10 00:00:00")
	to, _ := time.Parse(time.DateTime, "2023-10-14 00:00:00")

	period := domain.Period{
		From: from,
		To:   to,
	}

	successGeographyInfo := []domain.GeographyInfo{
		{
			Address:    "Address 1",
			Lat:        130.9,
			Lng:        -29.8,
			TotalPrice: 50,
		},
		{
			Address:    "Address 2",
			Lat:        28.1,
			Lng:        33.4,
			TotalPrice: 100,
		},
	}

	type fields struct {
		statsRepository StatsRepository
	}
	type args struct {
		ctx  context.Context
		from time.Time
		to   time.Time
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
			name: "repository error",
			fields: fields{
				statsRepository: statsRepositoryMock,
			},
			args: args{
				ctx:  ctx,
				from: from,
				to:   to,
			},
			initMocks: func() {
				statsRepositoryMock.EXPECT().GetGeographyByPeriod(gomock.Any(), gomock.Eq(period)).Return(nil, testingError)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				statsRepository: statsRepositoryMock,
			},
			args: args{
				ctx:  ctx,
				from: from,
				to:   to,
			},
			initMocks: func() {
				statsRepositoryMock.EXPECT().GetGeographyByPeriod(gomock.Any(), gomock.Eq(period)).Return(successGeographyInfo, nil)
			},
			want:    successGeographyInfo,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.initMocks()
			s := NewStats(tt.fields.statsRepository)
			got, err := s.GetGeographyByPeriod(tt.args.ctx, tt.args.from, tt.args.to)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}
