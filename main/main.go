package main

import (
	"alertBot/bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	//"main/test"
)

var botToken = os.Getenv("BOT_TOKEN")

func main() {
	if botToken == "" {
		log.Panic("BOT_TOKEN environment variable is not set")
	}
	botAPI, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	botAPI.Debug = true
	log.Printf("Authorized on account %s", botAPI.Self.UserName)

	go bot.ListenForUpdates(botAPI)
	//
	go bot.GoCheckForAlerts(botAPI)

	//telegram.GoNow()

	//telegram.Run()
	select {}
}
