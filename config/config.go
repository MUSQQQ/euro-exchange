package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ChecksNumber    int    `mapstructure:"X"`
	ChecksFrequency int    `mapstructure:"Y"`
	Timeout         int    `mapstructure:"TIMEOUT"`
	ExchangeURL     string `mapstructure:"EXCHANGE_URL"`

	LogFileName string `mapstructure:"LOG_FILE_NAME"`
}

func NewConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
	var cfg Config

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %v", err)
	}
	err := viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	return &cfg
}
