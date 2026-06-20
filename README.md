# Compogo Redis

Адаптер [Redis](https://redis.io/) для фреймворка [Compogo](https://github.com/Compogo/compogo).

На основе [go-redis](https://github.com/redis/go-redis) предоставляет:

* Клиент для работы с Redis
* Настройку через DSN (Data Source Name)
* Интеграцию с compogo-cache как драйвер

## Установка

```shell
go get github.com/Compogo/redis
```

## Быстрый старт

```go
package main

import (
    "context"
    "github.com/Compogo/compogo"
    compogoRedis "github.com/Compogo/redis"
	"github.com/redis/go-redis/v9"
)

func main() {
    app := compogo.NewApp("myapp",
        compogo.WithComponents(&compogoRedis.Component),
    )

    app.AddComponents(&compogo.Component{
        Name: "my_service",
        Init: compogo.StepFunc(func(container compogo.Container) error {
            return container.Invoke(func(client *redis.Client) error {
                ctx := context.Background()
                return client.Set(ctx, "key", "value", 0).Err()
            })
        }),
    })

    if err := app.Serve(); err != nil {
        panic(err)
    }
}
```

## Конфигурация через DSN

```shell
# Базовый формат
--cache.redis.dsn=redis://localhost:6379/0

# С аутентификацией
--cache.redis.dsn=redis://user:pass@localhost:6379/0

# С дополнительными параметрами
--cache.redis.dsn=redis://localhost:6379/0?dial_timeout=5s&pool_size=20&max_retries=3
```

## Интеграция с [Compogo Cache](https://github.com/Compogo/cache)

```go
import (
    "github.com/Compogo/redis/registration"
)

func main() {
    app := compogo.NewApp("myapp",
        compogo.WithComponents(&registration.Component),
    )
    // Теперь cache.NewCache() использует Redis
}
```

## Зависимости

* [Compogo](https://github.com/Compogo/compogo) — основной фреймворк
* [go-redis](https://github.com/redis/go-redis) — клиент Redis

## Лицензия

```plantuml
MIT License

Copyright (c) 2026 Compogo

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

```
