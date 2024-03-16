package greetings

import (
	"log"
	"mm-bot/internal/bot"
	"time"
)

func Main(username string) string {
	currentTime := time.Now()
	hour := currentTime.Hour()
	var greeting string

	if hour >= 7 && hour < 12 {
		greeting = "Доброе утро!"
	} else if hour >= 12 && hour < 18 {
		greeting = "Добрый день!"
	} else if hour >= 18 && hour < 22 {
		greeting = "Добрый вечер!"
	} else {
		greeting = "Доброй ночи!"
	}
	return greeting + " " + username
}

func BackupList(username string) {
	log.Printf("Пользователь %s запросил список баз\n", username)

	bot.B.SendMessage(username, "получаю список баз")
}
