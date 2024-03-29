package consumer

import (
	"context"
	"fmt"

	"github.com/nafisalfiani/ketson-notification-service/entity"
	"github.com/nafisalfiani/ketson-notification-service/lib/broker"
	"github.com/nafisalfiani/ketson-notification-service/lib/log"
	"github.com/nafisalfiani/ketson-notification-service/lib/parser"
	"github.com/nafisalfiani/ketson-notification-service/usecase/mailer"
)

type Interface interface {
	StartRegistrationConsumer(ctx context.Context)
	StartTransactionConsumer(ctx context.Context)
}

type consumer struct {
	logger log.Interface
	broker broker.Interface
	json   parser.JSONInterface
	mailer mailer.Interface
}

func Init(logger log.Interface, broker broker.Interface, json parser.JSONInterface, mailer mailer.Interface) Interface {
	return &consumer{
		logger: logger,
		broker: broker,
		json:   json,
		mailer: mailer,
	}
}

func (c *consumer) StartRegistrationConsumer(ctx context.Context) {
	topic := entity.TopicNewRegistration
	tag := "notification-service-new-registration-consumer"
	c.logger.Info(ctx, fmt.Sprintf("starting consumer for topic: %v & tag: %v", topic, tag))

	ch, err := c.broker.Channel()
	if err != nil {
		c.logger.Error(ctx, err)
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare(topic, false, false, false, false, nil)
	if err != nil {
		c.logger.Error(ctx, err)
	}

	msgs, err := ch.Consume(queue.Name, tag, true, false, false, false, nil)
	if err != nil {
		c.logger.Error(ctx, err)
	}

	for d := range msgs {
		var user entity.User
		if err := c.json.Unmarshal(d.Body, &user); err != nil {
			c.logger.Error(ctx, err)
			continue
		}

		c.logger.Info(ctx, "sending registration email")
		if err := c.mailer.SendRegistrationEmail(ctx, user); err != nil {
			c.logger.Error(ctx, err)
			d.Reject(true)
		}
	}
}

func (c *consumer) StartTransactionConsumer(ctx context.Context) {
	topic := entity.TopicNewTransaction
	tag := "notification-service-new-transaction-consumer"
	c.logger.Info(ctx, fmt.Sprintf("starting consumer for topic: %v & tag: %v", topic, tag))

	ch, err := c.broker.Channel()
	if err != nil {
		c.logger.Error(ctx, err)
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare(topic, false, false, false, false, nil)
	if err != nil {
		c.logger.Error(ctx, err)
	}

	msgs, err := ch.Consume(queue.Name, tag, true, false, false, false, nil)
	if err != nil {
		c.logger.Error(ctx, err)
	}

	for d := range msgs {
		var trx entity.Transaction
		if err := c.json.Unmarshal(d.Body, &trx); err != nil {
			c.logger.Error(ctx, err)
			continue
		}

		c.logger.Info(ctx, "sending transaction email")
		if err := c.mailer.SendTransactionEmail(ctx, trx); err != nil {
			c.logger.Error(ctx, err)
			d.Reject(true)
		}
	}
}
