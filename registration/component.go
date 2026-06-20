package registration

import (
	"github.com/Compogo/cache"
	"github.com/Compogo/compogo"
	"github.com/Compogo/redis"
	"github.com/eko/gocache/lib/v4/store"
	redisStore "github.com/eko/gocache/store/redis/v4"
	redisClient "github.com/redis/go-redis/v9"
)

// Component — компонент регистрации Redis драйвера для gocache.
// Регистрирует драйвер "redis" в системе кэширования.
//
// После подключения этого компонента, пакет cache сможет использовать
// Redis как бекенд для кэширования.
//
// Пример:
//
//	app.AddComponents(
//	    &registration.Component, // регистрация драйвера для cache
//	)
var Component = compogo.Component{
	Dependencies: compogo.Components{
		&redis.Component,
	},
}

// Регистрация драйвера "redis" в системе cache.
// Использует redis.Cmdable как источник данных.
func init() {
	cache.Registration("redis", func(container compogo.Container) (store.StoreInterface, error) {
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
