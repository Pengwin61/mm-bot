package main

import (
	"fmt"
	"mm-bot/internal/bot"
	"mm-bot/internal/core"
	"mm-bot/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	b := bot.NewBot()

	username := "pengwin"

	b.SendMessage(username, core.Greetings(username))

	// b.UploadFile("pengwin", "1112763.jpg", []byte("hello"))
	// core.Greetings()

	r := gin.Default()
	handlers.InitHandlers(r)

	// RUN
	err := r.Run(fmt.Sprint("0.0.0.0", ":", "8080"))
	if err != nil {
		panic(err)
	}
}
