package main

import (
	"fmt"
	"mm-bot/internal/bot"
	"mm-bot/internal/configs"
	sshclient "mm-bot/internal/ssh"
	"mm-bot/internal/web"

	"github.com/spf13/viper"
)

func main() {
	configs.InitConfigsViper()
	bot.InitBot()

	sshClient, err := sshclient.NewClient()
	if err != nil {
		fmt.Println("я не могу создать ssh подключение")
	}
	stdout := sshClient.ExecuteCmd("hostname", viper.GetString("database.host"))

	fmt.Println("STDOUT:", stdout)

	web.InitGin()
}
