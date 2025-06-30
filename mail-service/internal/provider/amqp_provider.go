package provider

import (
	"github.com/muammarahlnn/learnyscape-backend/mail-service/internal/config"
	"github.com/muammarahlnn/learnyscape-backend/mail-service/internal/mq"
	pmq "github.com/muammarahlnn/learnyscape-backend/pkg/mq"
)

func BootstrapAMQP(cfg *config.Config) []pmq.AMQPConsumer {
	return []pmq.AMQPConsumer{
		mq.NewSendVerificationConsumer(rabbitmq, mailer),
	}
}
