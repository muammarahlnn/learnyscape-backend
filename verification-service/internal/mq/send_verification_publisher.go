package mq

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/bytedance/sonic"
	"github.com/muammarahlnn/learnyscape-backend/pkg/mq"
	"github.com/muammarahlnn/learnyscape-backend/verification-service/internal/constant"
	"github.com/muammarahlnn/learnyscape-backend/verification-service/internal/log"
	amqp "github.com/rabbitmq/amqp091-go"
)

type SendVerificationPublisher struct {
	channel  *amqp.Channel
	exchange string
}

func NewSendVerificationPublisher(conn *amqp.Connection) mq.AMQPPublisher {
	exchange := constant.SendVerificationExchange
	ch, err := conn.Channel()
	if err != nil {
		log.Logger.Fatalf("failed to open a channel: %v", err)
	}

	if err := ch.ExchangeDeclare(exchange, constant.ExchangeTypeTopic, true, false, false, false, nil); err != nil {
		if amqpErr, ok := err.(*amqp.Error); ok && amqpErr.Code != amqp.PreconditionFailed {
			log.Logger.Fatalf("failed to declare an exchange: %v", err)
		}
	}

	return &SendVerificationPublisher{
		channel:  ch,
		exchange: exchange,
	}
}

func (p *SendVerificationPublisher) Publish(ctx context.Context, event mq.AMQPEvent) error {
	bytes, err := sonic.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}

	for i := 1; i <= constant.AMQPRetryLimit; i++ {
		err = p.channel.PublishWithContext(
			ctx,
			p.exchange,
			event.Key(),
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        bytes,
			},
		)
		if err == nil {
			log.Logger.Info("message published: %s", string(bytes))
			return nil
		}

		log.Logger.Errorf("failed to publish message: %v", err)

		delay := time.Duration(math.Pow(2, float64(i))) * constant.AMQPRetryDelay * time.Second
		time.Sleep(delay)
	}

	log.Logger.Errorf("message is aborted: %v", err)
	return err
}

func (p *SendVerificationPublisher) Exchange() string {
	return p.exchange
}
