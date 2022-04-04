package pingpong

import (
	"tgbot/internal/dbCon"
	"tgbot/internal/tgApi"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartPingpong() {
	bot := dbCon.Bot
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message

			tgApi.SendMessage(update.Message.Chat.ID, update.Message.Text)
		}
	}
}
