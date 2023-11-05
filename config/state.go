package config

import (
	"log/slog"

	"github.com/go-chi/jwtauth/v5"
	"github.com/redis/go-redis/v9"

	"github.com/hossein1376/kala/internal/data"
)

type Application struct {
	Config *Config
	Logger *slog.Logger
	Models *data.Models
	RDB    *redis.Client
}

type Config struct {
	Environment environment `json:"environment"`
	Port        string      `json:"port"`
	Cache       bool        `json:"cache"`

	DB     DB     `json:"db"`
	JWT    JWT    `json:"jwt"`
	Logger Logger `json:"logger"`
}

type JWT struct {
	Secret string           `json:"secret"`
	Expire string           `json:"expire"`
	Token  *jwtauth.JWTAuth `json:"-"`
}

type Logger struct {
	Level level `json:"level"`
}

type DB struct {
	Sql   Sql   `json:"sql"`
	Redis Redis `json:"redis"`
}

type Sql struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type Redis struct {
	Host     string `json:"host"`
	Password string `json:"password"`
	Port     string `json:"port"`
	DB       int    `json:"db"`
}
