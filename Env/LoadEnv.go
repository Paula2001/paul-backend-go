package Env

import (
	"github.com/joho/godotenv"
	"os"
)

func LoadEnv() {
	if os.Getenv("ENV") == "testing" {
		godotenv.Load("../.env.testing")
	} else {
		godotenv.Load(".env")
	}

}
