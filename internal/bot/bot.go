package bot

import (
	"mm-bot/internal/ewrap"

	"github.com/mattermost/mattermost-server/v6/model"
)

const (
	API_URL  = "http://192.168.140.185:8065"
	API_KEY  = "5xesykmmcfyqfg9wzkjt8jk74h"
	BOT_NAME = "mr.robot"
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

func (b *Bot) SendMessage(username string, msg string) {
	user, _, err := b.Client.GetUserByUsername(username, username)
	ewrap.Error(err)

	channel, _, err := b.Client.CreateDirectChannel(b.bot.Id, user.Id)
	ewrap.Error(err)

	post := &model.Post{}
	post.ChannelId = channel.Id
	post.Message = msg

	_, _, err = b.Client.CreatePost(post)
	ewrap.Error(err)

}

// func (b *Bot) UploadFile(username string, filename string) {

// 	user, _, err := b.Client.GetUserByUsername(username, username)
// 	ewrap.Error(err)

// 	channel, _, err := b.Client.CreateDirectChannel(b.bot.Id, user.Id)
// 	ewrap.Error(err)

// 	post := &model.Post{}
// 	post.ChannelId = channel.Id
// 	_, _, err = b.Client.UploadFile(, post.ChannelId, filename)
// 	ewrap.Error(err)

// }
