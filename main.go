package main

import (
	"bot/bot"
	"bot/environment"
	"log"
)

func main() {

	environment.Setup()

	bot.Start()

	log.Println("All good")

}
