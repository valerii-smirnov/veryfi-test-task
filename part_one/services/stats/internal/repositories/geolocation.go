package repositories

import (
	"context"
	"errors"

	"github.com/valerii-smirnov/veryfi-test-task/part_one/stats/internal/domain"

	"googlemaps.github.io/maps"
)

var NoGeolocationResultError = errors.New("geolocation not found")

type Geolocation struct {
	mapsClient *maps.Client
}

func NewGeolocation(mapsClient *maps.Client) *Geolocation {
	return &Geolocation{
		mapsClient: mapsClient,
	}
}

func (g Geolocation) GetLocation(ctx context.Context, address string) (*domain.Geography, error) {
	result, err := g.mapsClient.Geocode(ctx, &maps.GeocodingRequest{
		Address: address,
	})

	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, NoGeolocationResultError
	}

	return &domain.Geography{
		Address: address,
		Lat:     result[0].Geometry.Location.Lat,
		Lng:     result[0].Geometry.Location.Lng,
	}, nil
}
