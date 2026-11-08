package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort         string
	ServerReadTimeout  time.Duration
	ServerWriteTimeout time.Duration

	DBPath       string
	DBMigrations string

	UploadDir       string
	MaxUploadSizeMB int64

	LogLevel string
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{
		ServerPort:   getEnv("SERVER_PORT", "8080"),
		DBPath:       getEnv("DB_PATH", "./data/shop.db"),
		DBMigrations: getEnv("DB_MIGRATIONS", "./migrations"),
		UploadDir:    getEnv("UPLOAD_DIR", "./uploads"),
		LogLevel:     getEnv("LOG_LEVEL", "info"),
	}

	cfg.ServerReadTimeout = parseDuration(getEnv("SERVER_READ_TIMEOUT", "10s"))
	cfg.ServerWriteTimeout = parseDuration(getEnv("SERVER_WRITE_TIMEOUT", "10s"))

	maxMB, err := strconv.ParseInt(getEnv("MAX_UPLOAD_SIZE_MB", "10"), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid MAX_UPLOAD_SIZE_MB: %w", err)
	}
	cfg.MaxUploadSizeMB = maxMB

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return fallback
}

func parseDuration(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		return 10 * time.Second
	}
	return d
}