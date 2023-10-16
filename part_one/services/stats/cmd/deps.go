package main

import (
	"context"
	"fmt"

	"github.com/valerii-smirnov/veryfi-test-task/part_one/pkg/grpc/stubs/document"
	"github.com/valerii-smirnov/veryfi-test-task/part_one/stats/config"
	"github.com/valerii-smirnov/veryfi-test-task/part_one/stats/internal/repositories/models"
	"github.com/valerii-smirnov/veryfi-test-task/part_one/stats/internal/transport/rest"
	"github.com/valerii-smirnov/veryfi-test-task/part_one/stats/lib/event"

	"github.com/labstack/echo/v4"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"googlemaps.github.io/maps"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func newApplicationContext() (context.Context, context.CancelFunc) {
	return context.WithCancel(context.Background())
}

func newDocumentServiceClient(ctx context.Context, cfg config.Config) (document.DocumentServiceClient, error) {
	target := fmt.Sprintf("%s:%d", cfg.DocumentServiceHost, cfg.DocumentServicePort)
	conn, err := grpc.DialContext(ctx, target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return document.NewDocumentServiceClient(conn), nil
}

func newJetStreamClient(cfg config.Config) (jetstream.JetStream, error) {
	conn, err := nats.Connect(fmt.Sprintf("nats://%s:%d", cfg.JetStreamHost, cfg.JetStreamPort))
	if err != nil {
		return nil, err
	}

	js, err := jetstream.New(conn)
	if err != nil {
		return nil, err
	}

	return js, nil
}

func runConsumers(ctx context.Context, consumer *event.Consumer) error {
	if err := consumer.RunDocumentCreatedConsumer(ctx); err != nil {
		return err
	}

	if err := consumer.RunDocumentRemovedConsumer(ctx); err != nil {
		return err
	}

	return nil
}

func newGormConn(cfg config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.PostgresDSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func migrateDB(db *gorm.DB) (*gorm.DB, error) {
	if err := db.AutoMigrate(
		models.Receipt{},
		models.Item{},
		models.Geography{},
	); err != nil {
		return nil, err
	}

	return db, nil
}

func newGoogleMapsClient(cfg config.Config) (*maps.Client, error) {
	return maps.NewClient(maps.WithAPIKey(cfg.GoogleMapsAPIKey))
}

func runRESTServer(cfg config.Config, statsTransport *rest.Stats) error {
	s := echo.New()
	s.GET("/stats/total/tax", statsTransport.GetTotalTaxByPeriod)
	s.GET("/stats/total/discount", statsTransport.GetTotalDiscountByPeriod)
	s.GET("/stats/geography", statsTransport.GetGeographyByPeriod)

	logrus.WithField("port", cfg.RestServerPort).Info("running REST server ...")
	return s.Start(fmt.Sprintf(":%d", cfg.RestServerPort))
}
