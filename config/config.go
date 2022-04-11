package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var PrivateKey string
var Publickey string

func InitEnvironmentVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Println("ERROR:", err.Error(), ", reading from env")

	}
	PrivateKey = os.Getenv("PRIVATE_KEY_INTERNAL_CALL")
	Publickey = os.Getenv("PUBLIC_KEY_INTERNAL_CALL")
}
