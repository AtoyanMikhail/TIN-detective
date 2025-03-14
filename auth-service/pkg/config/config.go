package config

import (
    "os"
)

type Config struct {
    GRPCPort string
    DBHost   string
    DBPort   string
    DBUser   string
    DBPass   string
    DBName   string
}

func Load() *Config {
    return &Config{
        GRPCPort: getEnv("GRPC_PORT", "50051"),
        DBHost:   getEnv("DB_HOST", "localhost"),
        DBPort:   getEnv("DB_PORT", "5432"),
        DBUser:   getEnv("DB_USER", "postgres"),
        DBPass:   getEnv("DB_PASSWORD", "postgres"),
        DBName:   getEnv("DB_NAME", "auth"),
    }
}

func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}