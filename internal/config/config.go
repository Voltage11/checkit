package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
	Logger   LoggerConfig
}

type DatabaseConfig struct {
	Host           string `env:"DB_HOST" env-default:"localhost"`
	Port           int    `env:"DB_PORT" env-default:"5432"`
	User           string `env:"DB_USER" env-default:"postgres"`
	Password       string `env:"DB_PASSWORD" env-required:"true"`
	DBName         string `env:"DB_NAME" env-required:"true"`
	SSLMode        string `env:"DB_SSL_MODE" env-default:"disable"`
	MaxConnections int    `env:"DB_MAX_CONNS" env-default:"25"`
}

type ServerConfig struct {
	Address      string        `env:"SERVER_HOST" env-default:"127.0.0.1"`
	Port         string        `env:"SERVER_PORT" env-default:"8080"`
	ReadTimeout  time.Duration `env:"SERVER_READ_TIMEOUT_SEC" env-default:"10s"`
	WriteTimeout time.Duration `env:"SERVER_WRITE_TIMEOUT_SEC" env-default:"10s"`
	IdleTimeout  time.Duration `env:"SERVER_IDLE_TIMEOUT_SEC" env-default:"60s"`
}

type LoggerConfig struct {
	Level string `env:"LOG_LEVEL" env-default:"info"`
}

func New() (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadConfig(".env", &cfg); err != nil {
		// Если файла нет, пробуем читать просто из окружения
		if err := cleanenv.ReadEnv(&cfg); err != nil {
			return nil, fmt.Errorf("ошибка загрузки переменных окружения: %w", err)
		}
	}
	return &cfg, nil

}

func (c *Config) GetDSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		c.Database.User, c.Database.Password, c.Database.Host, c.Database.Port, c.Database.DBName, c.Database.SSLMode)
}
