package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	DatabasePort     string `env:"DB_PORT" envDefault:"5432"`
	DatabaseHost     string `env:"DB_HOST" envDefault:"127.0.0.1"`
	DatabasePassword string `env:"DB_PASSWORD" envDefault:"alpharius"`
	DatabaseUser     string `env:"DB_USER" envDefault:"gym_map"`
	DatabaseName     string `env:"DB_NAME" envDefault:"gym_map"`

	Env string `env:"ENV" envDefault:"dev"`
}

func (c Config) GetDSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.DatabaseUser,
		c.DatabasePassword,
		c.DatabaseHost,
		c.DatabasePort,
		c.DatabaseName)
}

func GetConfig() (c Config) {
	err := env.ParseWithOptions(&c, env.Options{
		Prefix: "APP_",
	})
	if err != nil {
		log.Fatal(err)
	}
	return
}
