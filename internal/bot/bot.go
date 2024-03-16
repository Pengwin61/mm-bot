package bot

import (
	"mm-bot/internal/ewrap"

	"github.com/mattermost/mattermost-server/v6/model"
)

const (
	API_URL       = "http://192.168.140.185:8065"
	API_KEY       = "5xesykmmcfyqfg9wzkjt8jk74h"
	BOT_NAME      = "mr.robot"
	SLASH_COMMAND = "cqhdt3nnfbrt7bmshrezrd5pge"
)

type Bot struct {
	Client *model.Client4
	bot    *model.User
}

func NewBot() *Bot {
	client := model.NewAPIv4Client(API_URL)
	client.SetToken(API_KEY)
	bot, _, err := client.GetUserByUsername(BOT_NAME, BOT_NAME)
	ewrap.Error(err)

	return &Bot{Client: client, bot: bot}
}
