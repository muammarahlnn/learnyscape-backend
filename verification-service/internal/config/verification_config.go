package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type VerificationConfig struct {
	AccountVerificationTokenDuration time.Duration `mapstructure:"ACCOUNT_VERIFICATION_TOKEN_DURATION"`
}

func initVerificationConfig() *VerificationConfig {
	adminConfig := &VerificationConfig{}
	if err := viper.Unmarshal(&adminConfig); err != nil {
		log.Fatalf("error mapping verification config: %v", err)
	}

	return adminConfig
}
