package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Database  MysqlConfig
	SecretKey string
}

type MysqlConfig struct {
	Username string
	Password string
	Name     string
	Host     string
	Port     int
}

func NewConfig(path string) *Config {
	viper.SetConfigFile(path + "/config.json")
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	conf := Config{}
	err := viper.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}
	return &conf
}
