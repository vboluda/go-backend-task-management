package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort     uint
	JWTSecret   string
	DatabaseURL string
}

func New() *Config {
	return &Config{}
}

func (c *Config) LoadEnv() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  No se pudo cargar .env, se usarán variables del entorno")
	}

	port, err := strconv.ParseUint(getEnv("APP_PORT", "8080"), 10, 64)

	if err != nil {
		log.Fatal("❌ Cannot parse APP PORT")
	}
	c.AppPort = uint(port)
	c.JWTSecret = getEnv("JWT_SECRET", "")
	c.DatabaseURL = getEnv("DATABASE_URL", "")
	return c
}

func (c *Config) Validate() *Config {
	if c.JWTSecret == "" {
		log.Fatal("❌ JWT_SECRET no definido en el entorno")
	}
	return c
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}

func (c *Config) String() string {
	str := "------------------------------------------------------\n"
	str += "              TASK MANAGER v0.1                       \n"
	str += "              Vicente Boluda Vias 2025                \n"
	str += "------------------------------------------------------\n"
	str += fmt.Sprintf("APP_PORT: %d \n", c.AppPort)
	str += fmt.Sprintf("JWT SECRET (LENGTH): %d \n", len(c.JWTSecret))
	str += fmt.Sprintf("DATABASE URL %s", c.DatabaseURL)

	return str
}
