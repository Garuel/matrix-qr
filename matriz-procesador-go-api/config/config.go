package config

import (
	"os"

	"github.com/joho/godotenv"
)


type ConfigStruct struct {
    PORT      string
    NODE_API_URL string
}

func Load() *ConfigStruct {
_ = godotenv.Load() 

    return &ConfigStruct{
        PORT:         getEnv("PORT", "3000"),
        NODE_API_URL: getEnv("NODE_API_URL", "http://localhost:7787/stats/calculate"),
    }
}

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}