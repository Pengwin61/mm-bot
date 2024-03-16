package main

import (
	"mm-bot/internal/bot"
	"mm-bot/internal/configs"
	"mm-bot/internal/web"
)

func main() {
	configs.InitConfigsViper()
	bot.InitBot()

	web.InitGin()
}
