package config

import (
	"log"
	"github.com/kelseyhightower/envconfig"
)
type Config struct {
	DbConfig DbConfig
}

type DbConfig struct {
	Host string `required:"true"`
	Port int `default:"5432"`
	User string `default:"postgres"`
	Password string `default:""`
	Name string `default:"mdsdb"`
}

func ReadConfig()(Config) {
	var dbConfig DbConfig
	err:= envconfig.Process("db", &dbConfig)
	if err != nil {
		log.Fatal(err)
	}
	return Config{DbConfig: dbConfig}
}
