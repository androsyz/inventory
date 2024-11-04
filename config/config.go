package config

import (
	"context"

	zlog "github.com/androsyz/inventory/internal/pkg/log"
	"github.com/spf13/viper"
)

type Config struct {
	Database Database `mapstructure:",squash"`
	Http     Http     `mapstructure:",squash"`
}

type Database struct {
	Host     string `mapstructure:"DB_HOST" default:"localhost"`
	Port     string `mapstructure:"DB_PORT" default:"5432"`
	User     string `mapstructure:"DB_USER" default:"postgres"`
	Name     string `mapstructure:"DB_NAME" default:"postgres"`
	Password string `mapstructure:"DB_PASSWORD" default:"postgres"`
}

type Http struct {
	Port string `mapstructure:"APP_PORT" default:"3000"`
	Host string `mapstructure:"APP_HOST" default:"localhost"`
}

func NewConfig() *Config {
	cfg := &Config{}

	zlog.Info(context.Background(), nil, "loading configuration")

	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		zlog.Fatal(context.Background(), nil, "failed load env file: "+err.Error())
	}

	if err := viper.Unmarshal(cfg); err != nil {
		zlog.Panic(context.Background(), nil, err.Error())
		panic(err)
	}

	zlog.Info(context.Background(), nil, "configuration loaded successfully")

	return cfg
}
