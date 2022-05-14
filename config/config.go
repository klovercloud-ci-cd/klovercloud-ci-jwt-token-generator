package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var PrivateKey string
var Publickey string
var RunMode string

func InitEnvironmentVariables() {
	RunMode = os.Getenv("RUN_MODE")
	log.Println("[INFO] Run mode: ", RunMode)
	if RunMode == "" {
		RunMode = string(DEVELOP)
	}

	if RunMode != string(PRODUCTION) {
		//Load .env file
		err := godotenv.Load()
		if err != nil {
			log.Println("ERROR:", err.Error())
			return
		}
	}
	log.Println("RUN MODE:", RunMode)
	PrivateKey = os.Getenv("PRIVATE_KEY_FOR_INTERNAL_CALL")
	Publickey = os.Getenv("PUBLIC_KEY_FOR_INTERNAL_CALL")
}

// ENVIRONMENT run environment
type ENVIRONMENT string

const (
	// PRODUCTION production environment
	PRODUCTION = ENVIRONMENT("PRODUCTION")
	// DEVELOP development environment
	DEVELOP = ENVIRONMENT("DEVELOP")
	// TEST test environment
	TEST = ENVIRONMENT("TEST")
)
