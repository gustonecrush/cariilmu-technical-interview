package app

import (
	"log"
	"os"

	"github.com/gustonecrush/cariilmu-technical-interview/app/controllers"
	"github.com/joho/godotenv"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func Run() {
	var server    = controllers.Server{}
	var appConfig = controllers.AppConfig{}
	var dbConfig  = controllers.DBConfig{}
	
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error on loading .env file")
	}

	appConfig.AppName = getEnv("APP_NAME", "API")
	appConfig.AppEnv  = getEnv("APP_ENV", "development")
	appConfig.AppPort = getEnv("APP_PORT", "8000")

	dbConfig.DBHost     = getEnv("DB_HOST", "localhost")
	dbConfig.DBUser     = getEnv("DB_USERNAME", "")
	dbConfig.DBPassword = getEnv("DB_PASSWORD", "")
	dbConfig.DBPort     = getEnv("DB_PORT", "3336")
	dbConfig.DBName     = getEnv("DB_NAME", "")
	dbConfig.DBDriver   = getEnv("DB_DRIVER", "mysql")

	server.Initialize(appConfig, dbConfig)
	server.Run(":" + appConfig.AppPort)
}