package provider

import (
	"github.com/jmoiron/sqlx"
	"github.com/muammarahlnn/learnyscape-backend/pkg/database"
	encryptutil "github.com/muammarahlnn/learnyscape-backend/pkg/util/encrypt"
	"github.com/muammarahlnn/user-service/internal/config"
	"github.com/muammarahlnn/user-service/internal/repository"
)

var (
	db           *sqlx.DB
	bcryptHasher encryptutil.Hasher
	dataStore    repository.DataStore
)

func BootstrapGlobal(cfg *config.Config) {
	db = database.NewPostgres((*database.PostgresOptions)(cfg.Postgres))
	bcryptHasher = encryptutil.NewBcryptHasher(cfg.App.BCryptCost)
	dataStore = repository.NewDataStore(db)
}
