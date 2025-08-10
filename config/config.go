package config

import (
	"article/model"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
}

func (c *Config) InitEnv() error {
	err := godotenv.Load("article.env")
	if err != nil {
		return err
	}
	return err
}

func (c *Config) CatchError(err error) {
	if err != nil {
		panic(any(err))
	}
}

func (c *Config) GetDBConfig() model.DBConfig {
	return model.DBConfig{
		DBName:   os.Getenv("DB_NAME"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PWD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
	}
}
