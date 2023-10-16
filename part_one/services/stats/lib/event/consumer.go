package event

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"

	"github.com/valerii-smirnov/veryfi-test-task/part_one/stats/config"

	"github.com/nats-io/nats.go/jetstream"
)

type CreatedChan chan uint

func NewCreatedChan() CreatedChan {
	return make(CreatedChan)
}

type RemovedChan chan uint

func NewRemovedChan() RemovedChan {
	return make(RemovedChan)
}

type Consumer struct {
	js     jetstream.JetStream
	config config.Config

	createdChan CreatedChan
	removedChan RemovedChan
}

func NewConsumer(
	js jetstream.JetStream,
	config config.Config,
	createdChan CreatedChan,
	removedChan RemovedChan,
) *Consumer {
	return &Consumer{
		js:          js,
		config:      config,
		createdChan: createdChan,
		removedChan: removedChan,
	}
}

func (c Consumer) RunDocumentCreatedConsumer(ctx context.Context) error {
	consumer, err := c.js.CreateOrUpdateConsumer(ctx, c.config.StreamConfig.DocumentStreamName, jetstream.ConsumerConfig{
		Durable:   c.config.StreamConfig.CreatedDocumentConsumerName,
		AckPolicy: jetstream.AckExplicitPolicy,
		FilterSubject: fmt.Sprintf("%s.%s",
			c.config.StreamConfig.DocumentStreamName,
			c.config.StreamConfig.CreateDocumentSubjectName,
		),
	})

	_, err = consumer.Consume(func(msg jetstream.Msg) {
		defer func() {
			if err := msg.Ack(); err != nil {
				logrus.WithContext(ctx).WithError(err).Error("message acknowledgement error")
			}
		}()

		var mData msgData
		if err := json.Unmarshal(msg.Data(), &mData); err != nil {
			logrus.WithContext(ctx).WithError(err).Error("unmarshalling message data error")
			return
		}

		c.createdChan <- mData.ID
	})

	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("consumer configuration error")
		return err
	}

	return nil
}

func (c Consumer) RunDocumentRemovedConsumer(ctx context.Context) error {
	consumer, err := c.js.CreateOrUpdateConsumer(ctx, c.config.StreamConfig.DocumentStreamName, jetstream.ConsumerConfig{
		Durable:   c.config.StreamConfig.RemovedDocumentConsumerName,
		AckPolicy: jetstream.AckExplicitPolicy,
		FilterSubject: fmt.Sprintf("%s.%s",
			c.config.StreamConfig.DocumentStreamName,
			c.config.StreamConfig.RemoveDocumentSubjectName,
		),
	})

	_, err = consumer.Consume(func(msg jetstream.Msg) {
		defer func() {
			if err := msg.Ack(); err != nil {
				logrus.WithContext(ctx).WithError(err).Error("message acknowledgement error")
			}
		}()

		var mData msgData
		if err := json.Unmarshal(msg.Data(), &mData); err != nil {
			logrus.WithContext(ctx).WithError(err).Error("unmarshalling message data error")
			return
		}

		c.removedChan <- mData.ID
	})

	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("consumer configuration error")
		return err
	}

	return nil
}

type msgData struct {
	ID uint `json:"id"`
}
