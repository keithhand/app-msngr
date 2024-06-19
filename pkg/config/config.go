package config

import (
	"log"

	"github.com/spf13/viper"
)

var config *Config

func New() *Config {
	if config == nil {
		config = &Config{
			Kubernetes: Kubernetes{},
			Helm:       Helm{},
		}
		setupViper()
		readConfig()
		config.updateConfig()
	}

	return config
}

func setupViper() {
	viper.SetConfigName("msngr")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
}

func readConfig() {
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
		} else {
			// Config file was found but another error was produced
			log.Panicf("fatal error reading config file: %s", err.Error())
		}
	}
}

func (cfg *Config) updateConfig() {
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Panicf("fatal error updating config file: %s", err.Error())
	}
}
