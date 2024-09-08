package config

import (
	"fmt"
	"log/slog"
	"os"
)

type Config struct {
	Host        string
	Port        string
	DbName      string
	LogLevel    slog.Leveler
	IsDebugMode bool
}

func (c Config) GetAddress() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

func NewConfig() *Config {
	lvl := &slog.LevelVar{}
	lvl.Set(slog.LevelInfo)

	cfg := &Config{
		Host:     "0.0.0.0",
		Port:     "5001",
		DbName:   "database.db",
		LogLevel: lvl,
	}

	if os.Getenv("DEBUG") == "1" {
		lvl.Set(slog.LevelDebug)
		cfg.IsDebugMode = true
	}
	return cfg
}
