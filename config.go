package main

import (
	"fmt"
	"os"
)

type Config struct {
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
	JWTSecret  string
}

var Envs = initConfig()

func initConfig() Config {
	return Config{
		Port:       getENV("PORT", "8080"),
		DBUser:     getENV("DB_USER", "root"),
		DBPassword: getENV("DB_PASSWORD", ""),
		DBAddress:  fmt.Sprintf("$s:$s", getENV("DB_HOST", "127.0.0.1"), getENV("DB_PORT", "3306")),
		DBName:     getENV("DB_NAME", "lets-go"),
		JWTSecret:  getENV("JWT_SECRET", "U81LohcwYBKow8XXhCqvgC4iJe3mjW"),
	}
}

func getENV(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
