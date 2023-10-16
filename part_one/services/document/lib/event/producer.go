package event

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/valerii-smirnov/veryfi-test-task/part_one/document/config"

	"github.com/nats-io/nats.go/jetstream"
)

type Producer struct {
	jetStream jetstream.JetStream
	cfg       config.Config
}

func NewProducer(jetStream jetstream.JetStream, cfg config.Config) *Producer {
	return &Producer{jetStream: jetStream, cfg: cfg}
}

func (p Producer) ProduceCreate(ctx context.Context, id uint) error {
	dm := documentMessage{ID: id}
	b, err := json.Marshal(dm)
	if err != nil {
		return err
	}

	subject := fmt.Sprintf("%s.%s",
		p.cfg.StreamConfig.DocumentStreamName,
		p.cfg.StreamConfig.CreateDocumentSubjectName,
	)

	if _, err := p.jetStream.Publish(ctx, subject, b); err != nil {
		return err
	}

	return nil
}

func (p Producer) ProduceRemove(ctx context.Context, id uint) error {
	dm := documentMessage{ID: id}
	b, err := json.Marshal(dm)
	if err != nil {
		return err
	}

	subject := fmt.Sprintf("%s.%s",
		p.cfg.StreamConfig.DocumentStreamName,
		p.cfg.StreamConfig.RemoveDocumentSubjectName,
	)

	if _, err := p.jetStream.Publish(ctx, subject, b); err != nil {
		return err
	}

	return nil
}

type documentMessage struct {
	ID uint `json:"id"`
}
