package provider

import (
	"github.com/jmoiron/sqlx"
	"github.com/muammarahlnn/learnyscape-backend/auth-service/internal/config"
	"github.com/muammarahlnn/learnyscape-backend/auth-service/internal/repository"
	"github.com/muammarahlnn/learnyscape-backend/pkg/database"
	encryptutil "github.com/muammarahlnn/learnyscape-backend/pkg/util/encrypt"
	jwtutil "github.com/muammarahlnn/learnyscape-backend/pkg/util/jwt"
)

var (
	db           *sqlx.DB
	bcryptHasher encryptutil.Hasher
	jwtUtil      jwtutil.JWTUtil
	dataStore    repository.DataStore
)

func BootstrapGlobal(cfg *config.Config) {
	db = database.NewPostgres((*database.PostgresOptions)(cfg.Postgres))
	bcryptHasher = encryptutil.NewBcryptHasher(cfg.App.BCryptCost)
	jwtUtil = jwtutil.NewJWTUtil()
	dataStore = repository.NewDataStore(db)
}
