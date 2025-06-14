package provider

import (
	"github.com/jmoiron/sqlx"
	"github.com/muammarahlnn/learnyscape-backend/admin-service/internal/config"
	"github.com/muammarahlnn/learnyscape-backend/admin-service/internal/repository"
	"github.com/muammarahlnn/learnyscape-backend/pkg/database"
)

var (
	db        *sqlx.DB
	dataStore repository.DataStore
)

func BootstrapGlobal(cfg *config.Config) {
	db = database.NewPostgres((*database.PostgresOptions)(cfg.Postgres))
	dataStore = repository.NewDataStore(db)
}
