package scraper

import (
	"fmt"
	"time"
)

type Config struct {
	ThreadCount int           `env:"THREAD_COUNT" env-default:"5"`
	Timeout     time.Duration `env:"TIMEOUT"      env-default:"5s"`
	RetryCount  int           `env:"RETRY_COUNT"  env-default:"3"`
}

func (c *Config) String() string {
	return fmt.Sprintf("{ThreadCount:%d Timeout:%s RetryCount:%d}", c.ThreadCount, c.Timeout, c.RetryCount)
}
