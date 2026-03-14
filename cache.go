package redis

import (
	"fmt"

	"github.com/Compogo/compogo/logger"
	"github.com/redis/go-redis/v9"
)

// NewCache creates a new Redis client instance.
// It initializes the underlying go-redis client with the configured host, port,
// authentication credentials, and timeout settings.
//
// The informer is used to log the configuration for debugging purposes.
// Returns a configured *redis.Client ready for use.
func NewCache(config *Config, informer logger.Informer) *redis.Client {
	informer.Infof("[cache.redis] host - %s, port - %d", config.Host, config.Port)
	informer.Infof("[cache.redis] readTimeout - %s", config.ReadTimeout)
	informer.Infof("[cache.redis] writeTimeout - %s", config.WriteTimeout)
	informer.Infof("[cache.redis] connectTimeout - %s", config.ConnectTimeout)

	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Username: config.User,
		Password: config.Password,

		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
		PoolTimeout:  config.ConnectTimeout,
	})
}
