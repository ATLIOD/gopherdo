package config

import (
	"os"
	"path/filepath"
)

type Config struct {
    Port     string
    DataDir  string
    DBPath   string
}

func NewConfig() *Config {
    dataDir := "./data"
    return &Config{
        Port:    getEnv("PORT", "8080"),
        DataDir: getEnv("DATA_DIR", dataDir),
        DBPath:  filepath.Join(dataDir, "gopherdo.db"),
    }
}

func getEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}
