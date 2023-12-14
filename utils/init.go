package init

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Token string
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Token = os.Getenv("TOKEN")
}
