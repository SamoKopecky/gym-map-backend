package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v11"
)

type Storage string

const (
	LOCAL Storage = "local"
	S3    Storage = "s3"
)

const CONFIG_PREFIX = "APP_"

type Config struct {
	DatabasePort     string `env:"DB_PORT" envDefault:"5432"`
	DatabaseHost     string `env:"DB_HOST" envDefault:"127.0.0.1"`
	DatabasePassword string `env:"DB_PASSWORD" envDefault:"alpharius"`
	DatabaseUser     string `env:"DB_USER" envDefault:"gym_map"`
	DatabaseName     string `env:"DB_NAME" envDefault:"gym_map"`

	Env string `env:"ENV" envDefault:"dev"`

	KeycloakBaseUrl           string `env:"KC_BASE_URL" envDefault:"http://localhost:8080"`
	KeycloakAdminClientId     string `env:"KC_ADMIN_CLIENT_ID" envDefault:"admin-cli"`
	KeycloakAdminClientSecret string `env:"KC_ADMIN_CLIENT_SECRET"`
	KeycloakRealm             string `env:"KC_REALM" envDefault:"gym-map"`

	StorageType         Storage `env:"STORAGE_TYPE" envDefault:"local"`
	StorageLocalPath    string  `env:"STORAGE_LOCAL_PATH" envDefault:"./files"`
	StorageS3Endpoint   string  `env:"STORAGE_S3_ENDPOINT"`
	StorageS3Region     string  `env:"STORAGE_S3_REGION"`
	StorageS3AccessKey  string  `env:"STORAGE_S3_ACCESS_KEY"`
	StorageS3SecretKey  string  `env:"STORAGE_S3_SECRET_KEY"`
	StorageS3BucketName string  `env:"STORAGE_S3_BUCKET_NAME" envDefault:"data"`
}

func (c Config) GetDSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.DatabaseUser,
		c.DatabasePassword,
		c.DatabaseHost,
		c.DatabasePort,
		c.DatabaseName,
	)
}

func GetConfig() (c Config) {
	err := env.ParseWithOptions(&c, env.Options{
		Prefix: CONFIG_PREFIX,
	})
	if err != nil {
		log.Fatal(err)
	}
	return
}
