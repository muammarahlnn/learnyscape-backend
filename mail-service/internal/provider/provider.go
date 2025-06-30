package provider

import (
	"github.com/muammarahlnn/learnyscape-backend/mail-service/internal/config"
	"github.com/muammarahlnn/learnyscape-backend/pkg/mq"
	smtputil "github.com/muammarahlnn/learnyscape-backend/pkg/util/smtp"
	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	rabbitmq *amqp.Connection
	mailer   smtputil.Mailer
)

func BootstrapGlobal(cfg *config.Config) {
	rabbitmq = mq.NewAMQP((*mq.AMQPOptions)(cfg.Amqp))
	mailer = smtputil.NewMailer()
}
