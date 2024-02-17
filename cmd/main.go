package main

import "mm-bot/internal/bot"

func main() {
	b := bot.NewBot()

	b.SendMessage("pengwin", "Hello!")

	// b.UploadFile("pengwin", "1112763.jpg", []byte("hello"))

}
