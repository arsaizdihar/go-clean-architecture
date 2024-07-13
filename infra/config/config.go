package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	DBFile    string
	JWTSecret string
	JWTExpire time.Duration
}

func Setup() *Config {
	godotenv.Load()

	return &Config{
		DBFile:    "db.sqlite3",
		JWTSecret: os.Getenv("JWT_SECRET"),
		JWTExpire: time.Minute * 10,
	}
}
