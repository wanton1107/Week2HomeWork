package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var AppConfig Config

type Datasource struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

type Config struct {
	Datasource Datasource `mapstructure:"datasource"`
}

func InitConfig() error {
	v := viper.New()
	v.SetConfigFile("application.yml")
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("read config file failed: %w", err)
	}

	active := v.GetString("profile.active")
	switch active {
	case "dev":
		v.SetConfigFile("application-dev.yml")
	case "prod":
		v.SetConfigFile("application-prod.yml")
	}

	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("read config file failed: %w", err)
	}

	if err := v.Unmarshal(&AppConfig); err != nil {
		return fmt.Errorf("unmarshal config failed: %w", err)
	}

	log.Println(AppConfig)
	return nil
}
