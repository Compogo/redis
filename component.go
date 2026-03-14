package redis

import (
	"github.com/Compogo/compogo/component"
	"github.com/Compogo/compogo/container"
	"github.com/Compogo/compogo/flag"
	"github.com/redis/go-redis/v9"
)

// Component is a ready-to-use Compogo component that provides a Redis client.
// It automatically:
//   - Registers Config and Client in the DI container
//   - Adds command-line flags for Redis configuration
//   - Configures the client during Configuration phase
//   - Provides the client as both *redis.Client and redis.Cmdable interface
//
// Usage:
//
//	compogo.WithComponents(
//	    redis.Component,
//	    // ... your service components
//	)
//
// Then in your service:
//
//	type Service struct {
//	    redis redis.Cmdable  // interface with all Redis commands
//	    // or
//	    client *redis.Client  // direct access
//	}
var Component = &component.Component{
	Init: component.StepFunc(func(container container.Container) error {
		return container.Provides(
			NewConfig,
			NewCache,
			func(cache *redis.Client) redis.Cmdable { return cache },
		)
	}),
	BindFlags: component.BindFlags(func(flagSet flag.FlagSet, container container.Container) error {
		return container.Invoke(func(config *Config) {
			flagSet.StringVar(&config.Host, HostFieldName, HostDefault, "redis connection host")
			flagSet.Uint16Var(&config.Port, PortFieldName, PortDefault, "redis connection port")
			flagSet.StringVar(&config.User, UserFieldName, "", "user for authorization in redis")
			flagSet.StringVar(&config.Password, PasswordFieldName, "", "password for authorization in redis")
			flagSet.DurationVar(&config.ReadTimeout, ReadTimeoutFieldName, ReadTimeoutDefault, "read timeout")
			flagSet.DurationVar(&config.WriteTimeout, WriteTimeoutFieldName, WriteTimeoutDefault, "write timeout")
			flagSet.DurationVar(
				&config.ConnectTimeout,
				ConnectTimeoutFieldName,
				ConnectTimeoutDefault,
				"The time the client waits for a connection if all connections are busy",
			)
		})
	}),
	Configuration: component.StepFunc(func(container container.Container) error {
		return container.Invoke(Configuration)
	}),
}
