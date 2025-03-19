package config

import (
	"errors"
	"flag"
	"fmt"
	"net"
	"slices"

	"github.com/ilyakaznacheev/cleanenv"
	"gopkg.in/yaml.v3"
)

type (
	Config struct {
		LoggerConfig LoggerConfig `yaml:"logger"`
		ServerConfig ServerConfig `yaml:"server"`
	}

	LoggerConfig struct {
		Format LogFormat `yaml:"format" env:"LOG_FORMAT" env-description:"log format: text,json"`
		Level  LogLevel  `yaml:"level"  env:"LOG_LEVEL"  env-description:"log level: debug,info,warn,error"`
	}

	LogLevel  string
	LogFormat string

	ServerConfig struct {
		Host ServerHost `yaml:"host" env:"HOST" env-description:"server host (0.0.0.0)"`
		Port ServerPort `yaml:"port" env:"PORT" env-description:"server port (>= 1024)"`
	}

	ServerPort int
	ServerHost string
)

const (
	LogFormatText LogFormat = "text"
	LogFormatJSON LogFormat = "json"

	LogLevelDebug LogLevel = "debug"
	LogLevelInfo  LogLevel = "info"
	LogLevelWarn  LogLevel = "warn"
	LogLevelError LogLevel = "error"
)

var (
	ErrInvalidLogLevel  = errors.New("invalid log level")
	ErrInvalidLogFormat = errors.New("invalid log format")

	ErrInvalidServerPort = errors.New("invalid server port")
	ErrInvalidServerHost = errors.New("invalid server host")
)

func Load() (*Config, error) {
	var config Config

	filepath := flag.String("c", "config.yml", "path to config file")

	flag.Usage = cleanenv.FUsage(flag.CommandLine.Output(), &config, nil, flag.Usage)
	flag.Parse()

	err := cleanenv.ReadConfig(*filepath, &config)
	if err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	err = config.Validate()
	if err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	return &config, nil
}

func (c *Config) Validate() error {
	if err := c.LoggerConfig.Validate(); err != nil {
		return err
	}

	if err := c.ServerConfig.Validate(); err != nil {
		return err
	}

	return nil
}

func (c *Config) String() string {
	return fmt.Sprintf("{Logger: %s, Server: %s}", &c.LoggerConfig, &c.ServerConfig)
}

func (l *LoggerConfig) Validate() error {
	if l.Format == "" {
		return ErrInvalidLogFormat
	}

	if l.Level == "" {
		return ErrInvalidLogLevel
	}

	return nil
}

func (l *LoggerConfig) String() string {
	return fmt.Sprintf("{Format: %s, Level: %s}", l.Format, l.Level)
}

func (s *ServerConfig) Validate() error {
	if s.Host == "" {
		return ErrInvalidServerHost
	}

	if s.Port == 0 {
		return ErrInvalidServerPort
	}

	return nil
}

func (s *ServerConfig) String() string {
	return fmt.Sprintf("{Host: %s, Port: %d}", s.Host, s.Port)
}

func (lf *LogFormat) UnmarshalYAML(node *yaml.Node) error {
	var value string

	err := node.Decode(&value)
	if err != nil {
		return fmt.Errorf("invalid log format: %w", err)
	}

	return lf.SetValue(value)
}

func (lf *LogFormat) SetValue(value string) error {
	formats := []LogFormat{LogFormatText, LogFormatJSON}

	if !slices.Contains(formats, LogFormat(value)) {
		return ErrInvalidLogFormat
	}

	*lf = LogFormat(value)

	return nil
}

func (ll *LogLevel) UnmarshalYAML(node *yaml.Node) error {
	var value string

	err := node.Decode(&value)
	if err != nil {
		return fmt.Errorf("invalid log level: %w", err)
	}

	return ll.SetValue(value)
}

func (ll *LogLevel) SetValue(value string) error {
	levels := []LogLevel{LogLevelDebug, LogLevelInfo, LogLevelWarn, LogLevelError}

	include := slices.Contains(levels, LogLevel(value))
	if !include {
		return ErrInvalidLogLevel
	}

	*ll = LogLevel(value)

	return nil
}

func (sp *ServerPort) UnmarshalYAML(node *yaml.Node) error {
	var value int

	err := node.Decode(&value)
	if err != nil {
		return fmt.Errorf("invalid server port: %w", err)
	}

	return sp.SetValue(value)
}

func (sp *ServerPort) SetValue(value int) error {
	if value < 1024 { //nolint:mnd
		return ErrInvalidServerPort
	}

	*sp = ServerPort(value)

	return nil
}

func (sh *ServerHost) UnmarshalYAML(node *yaml.Node) error {
	var value string

	err := node.Decode(&value)
	if err != nil {
		return fmt.Errorf("invalid server host: %w", err)
	}

	return sh.SetValue(value)
}

func (sh *ServerHost) SetValue(value string) error {
	ip := net.ParseIP(value)
	if ip == nil {
		return ErrInvalidServerHost
	}

	*sh = ServerHost(ip.String())

	return nil
}
