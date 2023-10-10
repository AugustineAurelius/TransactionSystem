package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Port           string `mapstructure:"PORT"`
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`
	REDISHOST      string `mapstructure:"REDISHOST"`
}

func LoadConfig() (config *Config, err error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("error while parse cfg :%w", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("error while unmarshal cfg :%w", err)
	}
	return
}
