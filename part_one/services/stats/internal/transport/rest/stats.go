package rest

import (
	"net/http"
	"time"

	"github.com/valerii-smirnov/veryfi-test-task/part_one/stats/internal/transport/rest/messages"

	"github.com/labstack/echo/v4"
)

type Stats struct {
	statsUseCase StatsUseCase
}

func NewStats(statsUseCase StatsUseCase) *Stats {
	return &Stats{
		statsUseCase: statsUseCase,
	}
}

func (s Stats) GetTotalTaxByPeriod(ctx echo.Context) error {
	from, to, err := parsePeriod(ctx)
	if err != nil {
		return err
	}

	totalTax, err := s.statsUseCase.TotalTaxByPeriod(ctx.Request().Context(), from, to)
	if err != nil {
		return err
	}

	msg := messages.TotalTax{
		Total: totalTax.Total,
	}

	return ctx.JSON(http.StatusOK, msg)
}

func (s Stats) GetTotalDiscountByPeriod(ctx echo.Context) error {
	from, to, err := parsePeriod(ctx)
	if err != nil {
		return err
	}

	totalDiscount, err := s.statsUseCase.TotalDiscountByPeriod(ctx.Request().Context(), from, to)
	if err != nil {
		return err
	}

	msg := messages.TotalDiscount{
		Total: totalDiscount.Total,
	}

	return ctx.JSON(http.StatusOK, msg)
}

func (s Stats) GetGeographyByPeriod(ctx echo.Context) error {
	from, to, err := parsePeriod(ctx)
	if err != nil {
		return err
	}

	geographyInfos, err := s.statsUseCase.GetGeographyByPeriod(ctx.Request().Context(), from, to)
	if err != nil {
		return err
	}

	geoMessages := make([]messages.Geography, 0, len(geographyInfos))
	for _, geographyInfo := range geographyInfos {
		geoMessages = append(geoMessages, messages.Geography{
			Address:    geographyInfo.Address,
			Lat:        geographyInfo.Lat,
			Lng:        geographyInfo.Lng,
			TotalPrice: geographyInfo.TotalPrice,
		})
	}

	return ctx.JSON(http.StatusOK, geoMessages)

}

func parsePeriod(ctx echo.Context) (time.Time, time.Time, error) {
	fromStr := ctx.QueryParam("from")
	from, err := time.Parse(time.DateTime, fromStr)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	toStr := ctx.QueryParam("to")
	to, err := time.Parse(time.DateTime, toStr)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	return from, to, nil
}
