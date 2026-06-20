package redis

import (
	"github.com/Compogo/compogo"
	"github.com/Compogo/compogo/flag"
	"github.com/redis/go-redis/v9"
)

// Component — компонент Redis для Compogo.
// Регистрирует конфигурацию и клиент в DI-контейнере.
//
// Пример подключения:
//
//	app.AddComponents(&redis.Component)
//
//	var client *redis.Client
//	container.Invoke(func(c *redis.Client) { client = c })
//	val, err := client.Get(ctx, "key").Result()
var Component = compogo.Component{
	Init: compogo.StepFunc(func(container compogo.Container) error {
		return container.Provides(
			NewConfig,
			NewCache,
			func(cache *redis.Client) redis.Cmdable { return cache },
		)
	}),
	BindFlags: compogo.BindFlags(func(flagSet flag.FlagSet, container compogo.Container) error {
		return container.Invoke(func(config *Config) {
			flagSet.StringVar(&config.DSN, DsnFieldName, DsnDefault, "connection string in Redis")
		})
	}),
	Configuration: compogo.StepFunc(func(container compogo.Container) error {
		return container.Invoke(Configuration)
	}),
}
