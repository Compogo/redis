package redis

import (
	"github.com/Compogo/compogo"
)

const (
	// DsnFieldName — имя поля для DSN (Data Source Name) строки подключения.
	DsnFieldName = "cache.redis.dsn"
)

// DsnDefault — DSN по умолчанию для локальной разработки.
// Формат: redis://[user:password@]host:port[/db][?options]
var DsnDefault = "redis://localhost:6789/0?dial_timeout=3&read_timeout=6s&max_retries=2"

// Config содержит конфигурацию Redis.
// Использует DSN строку для настройки всех параметров подключения.
//
// Поддерживаемые параметры в DSN:
//   - dial_timeout — таймаут подключения
//   - read_timeout — таймаут чтения
//   - write_timeout — таймаут записи
//   - pool_size — размер пула соединений
//   - max_retries — максимальное количество попыток
//   - и другие параметры библиотеки go-redis
//
// Примеры DSN:
//   - redis://localhost:6379/0
//   - redis://user:pass@redis-cluster:6379/0?dial_timeout=5s
//   - redis://localhost:6379?pool_size=10&max_retries=3
type Config struct {
	DSN string
}

// NewConfig создаёт новую конфигурацию.
func NewConfig() *Config {
	return &Config{}
}

// Configuration загружает конфигурацию из Configurator.
// Если DSN не задан, устанавливается значение по умолчанию.
func Configuration(config *Config, configurator compogo.Configurator) *Config {
	if config.DSN == "" || config.DSN == DsnDefault {
		configurator.SetDefault(DsnFieldName, DsnDefault)
		config.DSN = configurator.GetString(DsnFieldName)
	}

	return config
}
