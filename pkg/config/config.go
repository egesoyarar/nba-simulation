package config

import (
	"time"

	"github.com/spf13/viper"
)

func ReadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	config.GameDurationPerMin = config.GameDurationPerMin * time.Second
	config.MaxAttackDuration = config.MaxAttackDuration * time.Second

	return &config, nil
}
