package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Database DatabaseConfig
	Redis    RedisConfig
}

func NewConfig(file string) (*Config, error) {
	viper.SetConfigFile(file)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

type DatabaseConfig struct {
	Driver   string
	Username string
	Password string
	Host     string
	Port     int
	Database string
	Schema   string
}

func (d *DatabaseConfig) DSN() string {
	// "postgres://username:password@localhost:5432/database_name"
	// "postgres://user_dKwZY5:password_SxKEGb@pgsql.pi.local:5432/user_dKwZY5?sslmode=disable&search_path=public"
	return fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=disable&search_path=%s", d.Driver, d.Username, d.Password, d.Host, d.Port, d.Database, d.Schema)
}

type RedisConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	DB       string
}

func (r *RedisConfig) DSN() string {
	// "redis://<user>:<pass>@localhost:6379/<db>"
	return fmt.Sprintf("redis://%s:%s@%s:%s/%s", r.Username, r.Password, r.Host, r.Port, r.DB)
}
