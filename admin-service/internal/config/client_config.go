package config

import (
	"log"

	"github.com/spf13/viper"
)

type ClientConfig struct {
	UserURL string `mapstructure:"USER_URL"`
}

func initClientConfig() *ClientConfig {
	clientConfig := &ClientConfig{}
	if err := viper.Unmarshal(&clientConfig); err != nil {
		log.Fatalf("error mapping client config: %v", err)
	}

	return clientConfig
}
