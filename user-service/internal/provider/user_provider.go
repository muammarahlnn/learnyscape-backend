package provider

import (
	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
	pmq "github.com/muammarahlnn/learnyscape-backend/pkg/mq"

	"github.com/muammarahlnn/user-service/internal/config"
	"github.com/muammarahlnn/user-service/internal/constant"
	"github.com/muammarahlnn/user-service/internal/handler"
	"github.com/muammarahlnn/user-service/internal/mq"
	"github.com/muammarahlnn/user-service/internal/service"
)

func BootstrapUser(cfg *config.Config, router *gin.Engine) {
	kafkaAdmin.CreateTopic(
		constant.UserCreatedTopic,
		constant.UserCreatedTopicNumPartitions,
		constant.UserCreatedTopicReplicationFactor,
	)

	userCreatedProducer := mq.NewUserCreatedProducer(
		pmq.NewKafkaAsyncProducer(&pmq.KafkaProducerOptions{
			Brokers:        cfg.Kafka.Brokers,
			RetryMax:       cfg.Kafka.RetryMax,
			FlushFrequency: cfg.Kafka.FlushFrequency,
			ReturnSuccess:  cfg.Kafka.ReturnSuccess,
			Topic:          constant.UserCreatedTopic,
			Acks:           sarama.WaitForLocal,
		}),
	)

	userService := service.NewUserService(dataStore, bcryptHasher, userCreatedProducer)
	userHandler := handler.NewUserHanlder(userService)

	userHandler.Route(router)
}
