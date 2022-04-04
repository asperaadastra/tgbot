package dbCon

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	Bot *tgbotapi.BotAPI
)

func Connect() {
	connectTgBot()
}

func connectTgBot() {
	b, err := tgbotapi.NewBotAPI("5198414170:AAFPd4v3t0ty5cnSRJFN_Qjxp-e7twVk8aA")
	if err != nil {
		log.Panic(err)
	}

	b.Debug = true

	log.Printf("Authorized on account %s", b.Self.UserName)

	Bot = b
}
