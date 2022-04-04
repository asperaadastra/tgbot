package chatter

import (
	"fmt"
	"tgbot/internal/dbCon"
	"tgbot/internal/tgApi"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	lobby     []query
	chatRooms []chatRoom
)

type (
	query struct {
		Id       int64
		Name     string
		Password string
	}

	chatRoom struct {
		Members  []chatMember
		Password string
	}

	chatMember struct {
		Name string
		Id   int64
	}
)

func StartChatter() {
	bot := dbCon.Bot
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		updateHandler(update)
	}
}

func updateHandler(u tgbotapi.Update) {
	id := u.Message.From.ID
	name := u.Message.From.FirstName
	message := u.Message.Text

	if u.Message != nil {

		if message == "exit" {
			deleteFromChat(id)
			deleteFromLobby(id)
			return
		}

		chat := getChatRoom(id)
		if len(chat.Members) != 0 {
			for _, m := range chat.Members {
				if m.Id != id {
					tgApi.SendMessage(m.Id, fmt.Sprintf("[%s] %s", name, message))
				}
			}
			return
		}

		// if no chatroom found
		if newChat(name, id, message) {
			return
		}

		if inLobby(id) {
			return
		}

		lobby = append(lobby, query{id, name, message})
		tgApi.SendMessage(id, "You are in lobby now. Send your password to your friend to join you")

	}
}
