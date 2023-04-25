package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type dbConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

func getENV(key, defaultVal string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	env := os.Getenv(key)
	fmt.Println(env)
	if env == "" {
		return defaultVal
	}
	return env
}

var (
	ENV      = getENV("ENV", "testing") // testing as default to skip auth middleware during unit test
	AppName  = "sea-labs-library"
	DBConfig = dbConfig{
		Host:     getENV("HOST", ""),
		User:     getENV("DB_USER", ""),
		Password: getENV("DB_PASS", ""),
		DBName:   getENV("DB_NAME", ""),
		Port:     getENV("PORT", ""),
	}
)
