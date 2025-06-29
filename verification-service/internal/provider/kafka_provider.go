package provider

import (
	"github.com/IBM/sarama"
	pmq "github.com/muammarahlnn/learnyscape-backend/pkg/mq"
	"github.com/muammarahlnn/learnyscape-backend/verification-service/internal/config"
	"github.com/muammarahlnn/learnyscape-backend/verification-service/internal/constant"
	"github.com/muammarahlnn/learnyscape-backend/verification-service/internal/mq"
)

func BootstrapKafka(
	kafkaConfig *config.KafkaConfig,
	verificationConfig *config.VerificationConfig,
) []pmq.KafkaConsumer {
	sendVerificationPublisher := mq.NewSendVerificationPublisher(rabbitmq)

	return []pmq.KafkaConsumer{
		mq.NewUserCreatedConsumer(
			verificationConfig,
			pmq.NewKafkaConsumerGroup(&pmq.KafkaConsumerOptions{
				Brokers:       kafkaConfig.Brokers,
				ConsumerGroup: constant.UserCreatedConsumerGroup,
				InitialOffset: sarama.OffsetOldest,
			}),
			dataStore,
			sendVerificationPublisher,
		),
	}
}
