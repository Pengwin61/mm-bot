package bot

import (
	"mm-bot/internal/ewrap"

	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/spf13/viper"
)

type Bot struct {
	Client *model.Client4
	bot    *model.User
}

var B *Bot

func InitBot() {
	B = NewBot()
}

func NewBot() *Bot {
	client := model.NewAPIv4Client(viper.GetString("bot.apiurl"))
	client.SetToken(viper.GetString("bot.apikey"))
	bot, _, err := client.GetUserByUsername(viper.GetString("bot.name"), viper.GetString("bot.name"))
	ewrap.Error(err)

	return &Bot{Client: client, bot: bot}
}
