package main

import (
	"mm-bot/internal/bot"
	"mm-bot/internal/configs"
	"mm-bot/internal/core"
	"mm-bot/internal/web"
)

func main() {
	configs.InitConfigsViper()
	b := bot.NewBot()

	username := "pengwin"

	b.SendMessage(username, core.Greetings(username))

	// b.UploadFile("pengwin", "1112763.jpg", []byte("hello"))
	// core.Greetings()

	web.InitGin()
}
