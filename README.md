# Compogo Redis 🗄️

**Compogo Redis** — это полноценный Redis-клиент для Compogo, построенный поверх официальной библиотеки [go-redis](https://github.com/redis/go-redis). Предоставляет полный доступ ко всем командам Redis через интерфейс `redis.Cmdable`, настраивается через флаги и может использоваться как самостоятельный клиент или как драйвер для централизованной системы кэширования [Compogo Cache](https://github.com/Compogo/cache).

## 🚀 Установка

```bash
go get github.com/Compogo/redis
```

### 📦 Быстрый старт

```go
package main

import (
    "context"
    "github.com/Compogo/compogo"
    "github.com/Compogo/redis"
)

func main() {
    app := compogo.NewApp("myapp",
        compogo.WithOsSignalCloser(),
        redis.Component,  // добавляем Redis-клиент
        compogo.WithComponents(
            userServiceComponent,
        ),
    )

    if err := app.Serve(); err != nil {
        panic(err)
    }
}

// Использование в сервисе
var userServiceComponent = &component.Component{
    Dependencies: component.Components{redis.Component},
    Execute: component.StepFunc(func(c container.Container) error {
        return c.Invoke(func(redis redis.Cmdable) {
            service := &UserService{redis: redis}
            service.Start()
        })
    }),
}

type UserService struct {
    redis redis.Cmdable
}

func (s *UserService) GetUser(ctx context.Context, id int) (*User, error) {
    // Пытаемся достать из Redis
    data, err := s.redis.Get(ctx, fmt.Sprintf("user:%d", id)).Bytes()
    if err == nil {
        var user User
        json.Unmarshal(data, &user)
        return &user, nil
    }
    
    // Нет в кэше — грузим из БД
    user, err := s.db.LoadUser(id)
    if err != nil {
        return nil, err
    }
    
    // Кладём в Redis
    data, _ = json.Marshal(user)
    s.redis.Set(ctx, fmt.Sprintf("user:%d", id), data, 5*time.Minute)
    
    return user, nil
}
```

### ✨ Возможности

#### 🎯 Полный API Redis через redis.Cmdable

Интерфейс `Cmdable` содержит все команды Redis — более 200 методов:

```go
// Строки
redis.Set(ctx, "key", "value", time.Hour)
val, err := redis.Get(ctx, "key").Result()

// Хеши
redis.HSet(ctx, "user:1", "name", "Alice")
redis.HGetAll(ctx, "user:1").Result()

// Списки
redis.LPush(ctx, "queue", "task1")
redis.BRPop(ctx, 0, "queue").Result()

// Множества
redis.SAdd(ctx, "tags", "go", "redis")
redis.SMembers(ctx, "tags").Result()

// Pub/Sub
pubsub := redis.Subscribe(ctx, "channel")

// Транзакции
redis.Watch(ctx, func(tx *redis.Tx) error {
    // ...
}, "key")

// Скрипты
script := redis.NewScript("return redis.call('set', KEYS[1], ARGV[1])")
script.Run(ctx, redis, []string{"key"}, "value").Result()
```

### 🔌 Два способа использования

#### Как самостоятельный Redis-клиент:

```go
type Service struct {
    redis redis.Cmdable  // интерфейс со всеми командами
    // или
    client *redis.Client  // прямой доступ к клиенту
}
```

#### Как драйвер для централизованной системы кэширования:

```go
import (
    "github.com/Compogo/cache"
    "github.com/Compogo/redis/registration"
)

app := compogo.NewApp("myapp",
    cache.Component,
    redis.Component,
    registration.Component,  // регистрирует "redis" драйвер
)

// Теперь можно использовать через cache.CacheInterface[[]byte]
// с флагом --cache.driver=redis
```
