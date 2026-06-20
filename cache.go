package redis

import (
	"github.com/Compogo/compogo"
	"github.com/redis/go-redis/v9"
)

// NewCache создаёт новый клиент Redis из DSN.
// Парсит DSN и создаёт клиент с соответствующими настройками.
func NewCache(config *Config, logger compogo.Logger) (*redis.Client, error) {
	options, err := redis.ParseURL(config.DSN)
	if err != nil {
		return nil, err
	}

	logger = logger.GetLogger("cache").GetLogger("redis")
	logger.Infof("host - %s", options.Addr)
	logger.Infof("readTimeout - %s", options.ReadTimeout)
	logger.Infof("writeTimeout - %s", options.WriteTimeout)
	logger.Infof("poolTimeout - %s", options.PoolTimeout)

	return redis.NewClient(options), nil
}
