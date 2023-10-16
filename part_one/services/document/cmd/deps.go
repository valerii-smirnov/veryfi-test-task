package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/valerii-smirnov/veryfi-test-task/part_one/document/config"
	"github.com/valerii-smirnov/veryfi-test-task/part_one/document/internal/repositories/models"
	grpctransport "github.com/valerii-smirnov/veryfi-test-task/part_one/document/internal/transport/grpc"
	"github.com/valerii-smirnov/veryfi-test-task/part_one/pkg/grpc/stubs/document"

	"github.com/fsnotify/fsnotify"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/sirupsen/logrus"
	"github.com/veryfi/veryfi-go/veryfi"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func newFileWatcher(cfg config.Config) (*fsnotify.Watcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	if err := watcher.Add(cfg.FolderPath); err != nil {
		return nil, err
	}

	return watcher, nil
}

func newApplicationContext() (context.Context, context.CancelFunc) {
	return context.WithCancel(context.Background())
}

func newGormConn(cfg config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.PostgresDSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func newVeryfiClient(cfg config.Config) (*veryfi.Client, error) {
	client, err := veryfi.NewClientV8(&veryfi.Options{
		ClientID:     cfg.VerifyConfig.ClientID,
		ClientSecret: cfg.VerifyConfig.ClientSecret,
		Username:     cfg.VerifyConfig.Username,
		APIKey:       cfg.VerifyConfig.APIKey,
	})

	if err != nil {
		return nil, err
	}

	return client, nil
}

func migrateDB(db *gorm.DB) (*gorm.DB, error) {
	if err := db.AutoMigrate(models.Document{}); err != nil {
		return nil, err
	}

	return db, nil
}

func runGRPCServer(doc *grpctransport.Document, cfg config.Config) error {
	logrus.WithField("port", cfg.GRPCServerPort).Info("running gRPC server...")
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPCServerPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	document.RegisterDocumentServiceServer(grpcServer, doc)
	return grpcServer.Serve(ln)
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

func documentsStreamInitiator(ctx context.Context, js jetstream.JetStream, cfg config.Config) (jetstream.JetStream, error) {
	stream, err := js.Stream(ctx, cfg.StreamConfig.DocumentStreamName)
	if stream != nil {
		return js, nil
	}

	if _, err = js.CreateStream(ctx, jetstream.StreamConfig{
		Name: cfg.StreamConfig.DocumentStreamName,
		Subjects: []string{
			fmt.Sprintf("%s.%s", cfg.StreamConfig.DocumentStreamName, cfg.StreamConfig.CreateDocumentSubjectName),
			fmt.Sprintf("%s.%s", cfg.StreamConfig.DocumentStreamName, cfg.StreamConfig.RemoveDocumentSubjectName),
		},
	}); err != nil {
		return nil, err
	}

	return js, nil
}
