package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	ServerAddr           string
	Port                 string
	TimeAdditionMS       int
	TimeSubtractionMS    int
	TimeMultiplicationMS int
	TimeDivisionMS       int
	ComputingPower       int
}

func LoadConfig() *Config {
	cfg := &Config{
		ServerAddr:           getEnv("SERVER_ADDR", "localhost"),
		Port:                 getEnv("PORT", "8080"),
		TimeAdditionMS:       getEnvInt("TIME_ADDITION_MS", 2000),
		TimeSubtractionMS:    getEnvInt("TIME_SUBTRACTION_MS", 2000),
		TimeMultiplicationMS: getEnvInt("TIME_MULTIPLICATION_MS", 3000),
		TimeDivisionMS:       getEnvInt("TIME_DIVISIONS_MS", 3000),
		ComputingPower:       getEnvInt("COMPUTING_POWER", 1),
	}
	log.Printf("Config loaded: %+v", cfg)
	return cfg
}

func getEnv(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	return val
}

func getEnvInt(key string, defaultVal int) int {
	valStr := os.Getenv(key)
	if valStr != "" {
		if val, err := strconv.Atoi(valStr); err == nil {
			return val
		}
	}
	return defaultVal
}
