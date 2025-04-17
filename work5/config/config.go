package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		Log   Log
		Redis Redis
		PG    PG
		NATS  NATS
		HTTP  HTTP
	}

	Log struct {
		Level string `env:"LOG_LEVEL" env-default:"info"`
	}

	HTTP struct {
		Port         int           `env:"HTTP_PORT"          env-required:"true"`
		ReadTimeout  time.Duration `env:"HTTP_READ_TIMEOUT"  env-required:"true"`
		WriteTimeout time.Duration `env:"HTTP_WRITE_TIMEOUT" env-required:"true"`
	}

	PG struct {
		URL      string `env:"POSTGRES_URL"       env-required:"true"`
		PoolSize int    `env:"POSTGRES_POOL_SIZE" env-required:"true"`
	}

	Redis struct {
		URL          string        `env:"REDIS_URL"           env-required:"true"`
		PoolSize     int           `env:"REDIS_POOL_SIZE"     env-required:"true"`
		ReadTimeout  time.Duration `env:"REDIS_READ_TIMEOUT"  env-required:"true"`
		WriteTimeout time.Duration `env:"REDIS_WRITE_TIMEOUT" env-required:"true"`
	}

	NATS struct {
		URL           string        `env:"NATS_URL"            env-required:"true"`
		MaxReconnect  int           `env:"NATS_MAX_RECONNECT"  env-required:"true"`
		ReconnectWait time.Duration `env:"NATS_RECONNECT_WAIT" env-required:"true"`
		Timeout       time.Duration `env:"NATS_TIMEOUT"        env-required:"true"`
	}
)

func NewConfig() (*Config, error) {
	var conf Config

	// err := cleanenv.ReadEnv(&conf)
	err := cleanenv.ReadConfig(".env", &conf)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return &conf, nil
}

func (pg *PG) ToDSN() string {
	return fmt.Sprintf("%s?sslmode=disable&pool_max_conns=%d", pg.URL, pg.PoolSize)
}

func (rd *Redis) ToDSN() string {
	return fmt.Sprintf("%s?pool_size=%d&read_timeout=%s&write_timeout=%s", rd.URL, rd.PoolSize, rd.ReadTimeout, rd.WriteTimeout)
}
