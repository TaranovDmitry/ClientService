package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port             string
	DomainServiceURL string
}

func New() (Config, error) {
	if err := initConfig(); err != nil {
		return Config{}, fmt.Errorf("failed to init config: %w", err)
	}

	return Config{
		Port:             viper.GetString("port"),
		DomainServiceURL: viper.GetString("domainServiceURL"),
	}, nil
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
