package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	DBHost                  string
	DBPort                  string
	DBUser                  string
	DBPass                  string
	DBName                  string
	APIPort                 string
	DragonFlyPort           string
	RedisURL                string
	BrevoAPIKey             string
	JWTSecret               string
	Environment             string
	ClientID                string
	ClientSecret            string
	ServerURL               string
	ProjectID               string
	FirebaseCredentialsFile string
}

var Env *Environment

// getEnv func to get env value
func getEnv(key string, required bool) string {
	value, ok := os.LookupEnv(key)
	if !ok && required {
		log.Fatalf("Missing or invalid environment key: '%s'", key)
	}
	return value
}

func LoadEnvironment() {
	if Env == nil {
		Env = new(Environment)
	}
	Env.DBHost = getEnv("DB_HOST", true)
	Env.DBPort = getEnv("DB_PORT", true)
	Env.DBUser = getEnv("DB_USER", true)
	Env.DBPass = getEnv("DB_PASS", true)
	Env.DBName = getEnv("DB_NAME", true)
	Env.APIPort = getEnv("API_PORT", true)
	Env.BrevoAPIKey = getEnv("BREVO_API_KEY", true)
	Env.DragonFlyPort = getEnv("DRAGONFLY_PORT", true)
	Env.JWTSecret = getEnv("JWT_SECRET", true)
	Env.Environment = getEnv("ENV", true)
	Env.ClientID = getEnv("GOOGLE_CLIENT_ID", true)
	Env.ClientSecret = getEnv("GOOGLE_CLIENT_SECRET", true)
	Env.RedisURL = getEnv("REDIS_URL", true)
	Env.ServerURL = getEnv("SERVER_URL", true)
	Env.ProjectID = getEnv("PROJECT_ID", true)
	Env.FirebaseCredentialsFile = getEnv("FIREBASE_CREDENTIALS_FILE", true)
}

func LoadEnvironmentFile() {
	env := os.Getenv("ENV")

	if err := godotenv.Load(); err != nil {
		if env == "production" {
			log.Println("No .env file found, falling back to system environment variables.")
		} else {
			fmt.Printf("Error on load environment file: %s\n", err)
		}
	}

	LoadEnvironment()
}
