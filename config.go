package redis

import (
	"time"

	"github.com/Compogo/compogo/configurator"
)

const (
	HostFieldName           = "cache.redis.host"
	PortFieldName           = "cache.redis.port"
	UserFieldName           = "cache.redis.auth.user"
	PasswordFieldName       = "cache.redis.auth.password"
	ReadTimeoutFieldName    = "cache.redis.timeout.read"
	WriteTimeoutFieldName   = "cache.redis.timeout.write"
	ConnectTimeoutFieldName = "cache.redis.timeout.connect"

	HostDefault           = "localhost"
	PortDefault           = uint16(6379)
	ReadTimeoutDefault    = 300 * time.Millisecond
	WriteTimeoutDefault   = 300 * time.Millisecond
	ConnectTimeoutDefault = 500 * time.Millisecond
)

type Config struct {
	Host           string
	Port           uint16
	User           string
	Password       string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	ConnectTimeout time.Duration
}

func NewConfig() *Config {
	return &Config{}
}

func Configuration(config *Config, configurator configurator.Configurator) *Config {
	if config.Host == "" || config.Host == HostDefault {
		configurator.SetDefault(HostFieldName, HostDefault)
		config.Host = configurator.GetString(HostFieldName)
	}

	if config.Port == 0 || config.Port == PortDefault {
		configurator.SetDefault(PortFieldName, PortDefault)
		config.Port = configurator.GetUint16(PortFieldName)
	}

	if config.ReadTimeout == 0 || config.ReadTimeout == ReadTimeoutDefault {
		configurator.SetDefault(ReadTimeoutFieldName, ReadTimeoutDefault)
		config.ReadTimeout = configurator.GetDuration(ReadTimeoutFieldName)
	}

	if config.WriteTimeout == 0 || config.WriteTimeout == WriteTimeoutDefault {
		configurator.SetDefault(WriteTimeoutFieldName, WriteTimeoutDefault)
		config.WriteTimeout = configurator.GetDuration(WriteTimeoutFieldName)
	}

	if config.ConnectTimeout == 0 || config.ConnectTimeout == ConnectTimeoutDefault {
		configurator.SetDefault(ConnectTimeoutFieldName, ConnectTimeoutDefault)
		config.ConnectTimeout = configurator.GetDuration(ConnectTimeoutFieldName)
	}

	if config.User == "" {
		config.User = configurator.GetString(UserFieldName)
	}

	if config.Password == "" {
		config.Password = configurator.GetString(PasswordFieldName)
	}

	return config
}
