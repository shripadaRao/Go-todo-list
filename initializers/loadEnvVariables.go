package initializers

import (
	"github.com/joho/godotenv"
	"os"
)

func LoadEnvVariables(key string) string {
	godotenv.Load(".env")
	
  	return os.Getenv(key)
}