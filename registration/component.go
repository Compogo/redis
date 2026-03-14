package registration

import (
	"github.com/Compogo/cache"
	"github.com/Compogo/compogo/component"
	"github.com/Compogo/compogo/container"
	"github.com/Compogo/redis"
	"github.com/eko/gocache/lib/v4/store"
	redisStore "github.com/eko/gocache/store/redis/v4"
	redisClient "github.com/redis/go-redis/v9"
)

// Component is a ready-to-use Compogo component that registers the Redis
// cache driver with the central cache system.
//
// It depends on redis.Component to ensure the underlying Redis client
// is initialized. The actual registration happens in init(), which is safe because
// cache.Registration only stores factory functions without requiring runtime state.
var Component = &component.Component{
	Dependencies: component.Components{
		redis.Component,
	},
}

// init registers the "redis" cache driver with the central cache system.
// The registration happens at program startup, independent of component lifecycle.
//
// The factory function receives a container and:
//   - Extracts the cache configuration (*cache.Config) and the Redis client
//   - Creates a redis.Store (compatible with store.StoreInterface)
//   - Returns it to the cache system for wrapping with metrics
//
// This driver can be selected with --cache.driver=redis.
func init() {
	cache.Registration("redis", func(container container.Container) (store.StoreInterface, error) {
		var cacheStore store.StoreInterface
		var err error

		err = container.Invoke(func(config *cache.Config, client redisClient.Cmdable) {
			cacheStore = redisStore.NewRedis(client, store.WithExpiration(config.Expiration))
		})

		if err != nil {
			return nil, err
		}

		return cacheStore, nil
	})
}
