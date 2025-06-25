package provider

import (
	"github.com/IBM/sarama"
	"github.com/muammarahlnn/learnyscape-backend/auth-service/internal/config"
	"github.com/muammarahlnn/learnyscape-backend/auth-service/internal/constant"
	"github.com/muammarahlnn/learnyscape-backend/auth-service/internal/mq"
	pmq "github.com/muammarahlnn/learnyscape-backend/pkg/mq"
)

func BootstrapKafka(cfg *config.KafkaConfig) []pmq.KafkaConsumer {
	return []pmq.KafkaConsumer{
		mq.NewUserCreatedConsumer(
			pmq.NewKafkaConsumerGroup(&pmq.KafkaConsumerOptions{
				Brokers:       cfg.Brokers,
				ConsumerGroup: constant.UserCreatedConsumerGroup,
				InitialOffset: sarama.OffsetOldest,
			}),
			dataStore,
		),
	}
}
