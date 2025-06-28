package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	App          *AppConfig
	HttpServer   *HttpServerConfig
	Postgres     *PostgresConfig
	Logger       *LoggerConfig
	Kafka        *KafkaConfig
	Verification *VerificationConfig
}

func InitConfig() *Config {
	configPath := parseConfigPath()
	viper.AddConfigPath(configPath)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error reading config file: %v", err)
	}

	return &Config{
		App:          initAppConfig(),
		HttpServer:   initHttpServerConfig(),
		Postgres:     initPostgresConfig(),
		Logger:       initLoggerConfig(),
		Kafka:        initKafkaConfig(),
		Verification: initVerificationConfig(),
	}
}

func parseConfigPath() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return filepath.Join(wd)
}
