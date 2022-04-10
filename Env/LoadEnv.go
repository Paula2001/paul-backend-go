package Env

import (
	"github.com/joho/godotenv"
	"os"
)

func LoadEnv() {
	if os.Getenv("ENVIRONMENT") == "testing" {
		godotenv.Load("../.env.testing")
	} else {
		godotenv.Load(".env")
	}

}
