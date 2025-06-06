package config

import (
	"log"

	"github.com/spf13/viper"
)

type PostgresConfig struct {
	Host            string `mapstructure:"DB_HOST"`
	Name            string `mapstructure:"DB_NAME"`
	Username        string `mapstructure:"DB_USER"`
	Password        string `mapstructure:"DB_PASSWORD"`
	SSLMode         string `mapstructure:"DB_SSL_MODE"`
	Port            string `mapstructure:"DB_PORT"`
	MaxIdleConn     string `mapstructure:"DB_MAX_IDLE_CONN"`
	MaxOpenConn     string `mapstructure:"DB_MAX_OPEN_CONN"`
	MaxConnLifetime string `mapstructure:"DB_CONN_MAX_LIFETIME"`
}

func initPostgresConfig() *PostgresConfig {
	redisConfig := &PostgresConfig{}
	if err := viper.Unmarshal(&redisConfig); err != nil {
		log.Fatalf("error mapping database config: %v", err)
	}

	return redisConfig
}
