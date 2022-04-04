package chatter

import (
	"fmt"
	"tgbot/internal/tgApi"
)

func getChatRoom(id int64) chatRoom {
	for _, ch := range chatRooms {
		for _, m := range ch.Members {
			if m.Id == id {
				return ch
			}
		}
	}
	return chatRoom{}
}

func newChat(name string, id int64, p string) bool {
	for k, l := range lobby {
		if l.Password == p {
			chat := chatRoom{
				Password: p,
			}
			chat.Members = append(chat.Members, chatMember{name, id})
			chat.Members = append(chat.Members, chatMember{l.Name, l.Id})
			chatRooms = append(chatRooms, chat)

			tgApi.SendMessage(id, fmt.Sprintf("You are in chat with %s", l.Name))
			tgApi.SendMessage(l.Id, fmt.Sprintf("%s have entered your chatroom", name))
			lobby = append(lobby[:k], lobby[k+1:]...)
			return true
		}

	}
	for k, ch := range chatRooms {
		if ch.Password == p {
			chatMates := ""
			for _, i := range ch.Members {
				tgApi.SendMessage(i.Id, fmt.Sprintf("%s joined your chatroom", name))
				chatMates += i.Name + ", "
			}
			tgApi.SendMessage(id, "You have entered chat with "+chatMates)
			chatRooms[k].Members = append(ch.Members, chatMember{name, id})
			return true
		}
	}
	return false
}

func deleteFromChat(id int64) {
	for i, ch := range chatRooms {
		for k, m := range ch.Members {
			if m.Id == id {
				chatRooms[i].Members = append(ch.Members[:k], ch.Members[k+1:]...)
				for _, r := range chatRooms[i].Members {
					tgApi.SendMessage(r.Id, fmt.Sprintf("%s quited chatroom", m.Name))
				}
				tgApi.SendMessage(m.Id, "You have quited the chat")
				return
			}
		}
	}
}

func deleteFromLobby(id int64) {
	for k, l := range lobby {
		if l.Id == id {
			lobby = append(lobby[:k], lobby[k+1:]...)
			tgApi.SendMessage(id, "You have quited lobby")
		}
	}
}

func inLobby(id int64) bool {
	for _, l := range lobby {
		if id == l.Id {
			tgApi.SendMessage(id, "You are already in lobby. type exit to set a new password for your chat")
			return true
		}

	}
	return false
}
