package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type DBConfig struct {
	Host     string `env:"DB_HOST" env-default:"localhost"`
	Port     string `env:"DB_PORT" env-default:"5432"`
	User     string `env:"DB_USER" env-default:"postgres"`
	Password string `env:"DB_PASSWORD" env-default:"postgres"`
}

func (c *DBConfig) DSN() string {
	return "host=" + c.Host + " port=" + string(c.Port) + " user=" + c.User + " dbname=postgres password=" + c.Password + " sslmode=disable"
}

type Config struct {
	// DB is the database configuration.
	DB DBConfig
}

var (
	cfg  Config
	once sync.Once
)

func Get() *Config {
	once.Do(func() {
		err := cleanenv.ReadEnv(&cfg)
		if err != nil {
			log.Fatal(err)
		}
	})
	return &cfg
}
