package provider

import (
	"github.com/jmoiron/sqlx"
	"github.com/muammarahlnn/learnyscape-backend/pkg/database"
	"github.com/muammarahlnn/learnyscape-backend/verification-service/internal/config"
	"github.com/muammarahlnn/learnyscape-backend/verification-service/internal/repository"
)

var (
	db        *sqlx.DB
	dataStore repository.DataStore
)

func BootstrapGlobal(cfg *config.Config) {
	db = database.NewPostgres((*database.PostgresOptions)(cfg.Postgres))
	dataStore = repository.NewDataStore(db)
}
