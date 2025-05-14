package config

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	Database DatabaseConfig `mapstructure:"database"`
	Redis    RedisConfig    `mapstructure:"redis"`
	JWT      JWTConfig      `mapstructure:"jwt"`
	Email    EmailConfig    `mapstructure:"email"`
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
	Driver   string `mapstructure:"driver"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Database string `mapstructure:"database"`
	Schema   string `mapstructure:"schema"`
}

func (d *DatabaseConfig) DSN() string {
	// "postgres://username:password@localhost:5432/database_name"
	// "postgres://user_dKwZY5:password_SxKEGb@pgsql.pi.local:5432/user_dKwZY5?sslmode=disable&search_path=public"
	return fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=disable&search_path=%s", d.Driver, d.Username, d.Password, d.Host, d.Port, d.Database, d.Schema)
}

type RedisConfig struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	DB       string `mapstructure:"db"`
}

func (r *RedisConfig) DSN() string {
	// "redis://<user>:<pass>@localhost:6379/<db>"
	return fmt.Sprintf("redis://%s:%s@%s:%d/%s", r.Username, r.Password, r.Host, r.Port, r.DB)
}

type JWTConfig struct {
	Secret   string        `mapstructure:"secret"`
	Duration time.Duration `mapstructure:"duration"`
}

type EmailConfig struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Subject  string `mapstructure:"subject"`
}

func (cfg *EmailConfig) DSN() string {
	return fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
}
