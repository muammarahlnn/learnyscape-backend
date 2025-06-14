package config

import (
	"log"

	"github.com/spf13/viper"
)

type ServiceConfig struct {
	AuthURL  string `mapstructure:"AUTH_SERVICE_URL"`
	AdminURL string `mapstructure:"ADMIN_SERVICE_URL"`
}

func initServiceConfig() *ServiceConfig {
	serviceConfig := &ServiceConfig{}
	if err := viper.Unmarshal(serviceConfig); err != nil {
		log.Fatalf("error mapping service config: %v", err)
	}

	return serviceConfig
}
