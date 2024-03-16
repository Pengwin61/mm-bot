package handlers

import (
	"fmt"
	"mm-bot/internal/bot"
	"mm-bot/internal/greetings"
)

func validate(cmd string, username string) {
	newOutput := fmt.Sprintln("проверь команду", "ты ввел:", cmd)

	switch cmd {
	case "":
		bot.B.SendMessage(username, greetings.Main(username))
	case "list":
		greetings.BackupList(username)
	case "image":
		// bot.B.UploadFile(username, "./upload/1112763.jpg", "robot.jpg")
		bot.B.SendMessage(username, "эта функция пока еще не работает")
	default:
		bot.B.SendMessage(username, newOutput)
	}

}
