package tgApi

import (
	"tgbot/internal/dbCon"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendMessage(id int64, message string) {
	bot := dbCon.Bot
	msg := tgbotapi.NewMessage(id, message)

	bot.Send(msg)
}
