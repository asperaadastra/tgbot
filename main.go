package main

import (
	"tgbot/internal/chatter"
	"tgbot/internal/dbCon"
	"tgbot/internal/pingpong"
)

func main() {
	dbCon.Connect()

	pingpong.StartPingpong()
	chatter.StartChatter()
}
