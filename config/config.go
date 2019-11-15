package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

//Config ...
type Config struct {
	Port         string
	DbStringConn string
	DbName       string
}

var config Config

//Init ...
func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file ", err)
	}

	config = Config{
		Port:         os.Getenv("PORT"),
		DbStringConn: os.Getenv("DB_STRING_CONNECTION"),
		DbName:       os.Getenv("DB_NAME"),
	}
}

//GetConfig ...
func GetConfig() *Config {
	return &config
}
