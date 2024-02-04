package environment

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	BotToken  string
	BotPrefix string

	ApiKey string
)

func Setup() {

	// Reading .env file
	godotenv.Load()

	// Getting necessary information from .env file
	BotToken = os.Getenv("BOT_TOKEN")
	log.Println("The bot token is " + BotToken)

	BotPrefix = os.Getenv("BOT_PREFIX")
	log.Println("The bot prefix is " + BotPrefix)

	// ApiKey is for gpt, which is currently not developed
	ApiKey = os.Getenv("API_KEY")
}
