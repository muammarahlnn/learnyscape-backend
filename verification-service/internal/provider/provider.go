package provider

import (
	"github.com/jmoiron/sqlx"
	"github.com/muammarahlnn/learnyscape-backend/pkg/database"
	"github.com/muammarahlnn/learnyscape-backend/pkg/mq"
	"github.com/muammarahlnn/learnyscape-backend/verification-service/internal/config"
	"github.com/muammarahlnn/learnyscape-backend/verification-service/internal/repository"
	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	db        *sqlx.DB
	rabbitmq  *amqp.Connection
	dataStore repository.DataStore
)

func BootstrapGlobal(cfg *config.Config) {
	db = database.NewPostgres((*database.PostgresOptions)(cfg.Postgres))
	rabbitmq = mq.NewAMQP((*mq.AMQPOptions)(cfg.Amqp))
	dataStore = repository.NewDataStore(db)
}
