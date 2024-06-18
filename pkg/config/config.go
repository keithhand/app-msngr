package config

import (
	"log"

	"github.com/spf13/viper"
)

func setupViper() {
	viper.SetConfigName("msngr")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
}

func Read() {
	setupViper()
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
		} else {
			// Config file was found but another error was produced
			log.Panicf("fatal error config file: %s", err.Error())
		}
	}
}
