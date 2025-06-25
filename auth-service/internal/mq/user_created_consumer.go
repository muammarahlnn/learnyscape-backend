package mq

import (
	"context"
	"math"
	"sync"
	"time"

	"github.com/IBM/sarama"
	"github.com/bytedance/sonic"
	"github.com/muammarahlnn/learnyscape-backend/auth-service/internal/constant"
	"github.com/muammarahlnn/learnyscape-backend/auth-service/internal/dto"
	"github.com/muammarahlnn/learnyscape-backend/auth-service/internal/entity"
	"github.com/muammarahlnn/learnyscape-backend/auth-service/internal/log"
	"github.com/muammarahlnn/learnyscape-backend/auth-service/internal/repository"
	"github.com/muammarahlnn/learnyscape-backend/pkg/mq"
)

type UserCreatedConsumer struct {
	consumer  sarama.ConsumerGroup
	dataStore repository.DataStore
	topic     string
	wg        *sync.WaitGroup
}

func NewUserCreatedConsumer(consumer sarama.ConsumerGroup, dataStore repository.DataStore) mq.KafkaConsumer {
	return &UserCreatedConsumer{
		consumer:  consumer,
		dataStore: dataStore,
		topic:     constant.UserCreatedTopic,
		wg:        &sync.WaitGroup{},
	}
}

func (c *UserCreatedConsumer) Consume(ctx context.Context) error {
	c.wg.Add(1)
	return c.consumer.Consume(ctx, []string{c.topic}, c)
}

func (c *UserCreatedConsumer) Handler() mq.KafkaHandler {
	return func(ctx context.Context, body []byte) error {
		event := &dto.UserCreatedEvent{}
		if err := sonic.Unmarshal(body, event); err != nil {
			return err
		}

		return c.dataStore.Atomic(ctx, func(ds repository.DataStore) error {
			// TODO: handle upsert for idempotency
			if err := ds.UserRepository().Create(ctx, &entity.CreateUserParams{
				ID:           event.Id,
				Username:     event.Username,
				Email:        event.Email,
				HashPassword: event.HashPassword,
			}); err != nil {
				return err
			}

			return nil
		})
	}
}

func (c *UserCreatedConsumer) Close() error {
	log.Logger.Infof("Closing consume for topic: %s", c.Topic())
	c.wg.Wait()

	return c.consumer.Close()
}

func (c *UserCreatedConsumer) Topic() string {
	return c.topic
}

func (c *UserCreatedConsumer) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (c *UserCreatedConsumer) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (c *UserCreatedConsumer) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	defer c.wg.Done()

	for message := range claim.Messages() {
		log.Logger.Infof("topic %v received a message %v", c.Topic(), string(message.Value))

		for i := 1; i < constant.KafkaConsumerRetryLimit+1; i++ {
			if err := c.Handler()(sess.Context(), message.Value); err != nil {
				log.Logger.Errorf("failed to consume message: %s", err)

				if i > constant.KafkaConsumerRetryLimit {
					log.Logger.Errorf("failed to consume message after %d retries: %s", constant.KafkaConsumerRetryLimit, err)
				} else {
					delay := math.Pow(constant.KafkaConsumerRetryDelay, float64(i))
					time.Sleep(time.Duration(delay) * time.Second)
					log.Logger.Infof("retrying to consume message, attempt %d", i)
				}
			} else {
				sess.MarkMessage(message, "")
				break
			}
		}

	}

	return nil
}
