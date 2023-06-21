package config

import (
	"os"
	"strconv"
	"strings"
)

type ApplicationConfig struct {
	Redis              RedisConfig
	CopyImage          string
	SearchImage        string
	ServiceAccountName string
	Namespace          string
}

type NewrelicConfig struct {
	Name    string
	License string
	Enabled bool
}

type RedisConfig struct {
	MasterName string
	Host       []string
	Password   string
	Database   int
}

type KafkaConfig struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Host        string   `json:"host"`
	UserName    *string  `json:"userName"`
	Password    *string  `json:"password"`
	Certificate []string `json:"certificate"`
}

func NewApplicationConfig() (*ApplicationConfig, error) {
	configuration := ApplicationConfig{}
	configuration.Namespace = os.Getenv("NAMESPACE")
	configuration.SearchImage = os.Getenv("SEARCH_IMAGE")
	configuration.CopyImage = os.Getenv("COPY_IMAGE")
	configuration.ServiceAccountName = os.Getenv("SERVICE_ACCOUNT_NAME")

	configuration.Redis.Host = strings.Split(os.Getenv("REDIS_HOST"), ",")
	configuration.Redis.Password = os.Getenv("REDIS_PASSWORD")
	configuration.Redis.MasterName = os.Getenv("REDIS_MASTERNAME")
	db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	configuration.Redis.Database = db

	return &configuration, nil
}
