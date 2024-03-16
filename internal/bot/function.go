package bot

import (
	"mm-bot/internal/core"
	"mm-bot/internal/ewrap"

	"github.com/mattermost/mattermost-server/v6/model"
)

func (b *Bot) SendMessage(username string, msg string) {
	user, _, err := b.Client.GetUserByUsername(username, username)
	ewrap.Error(err)

	channel, _, err := b.Client.CreateDirectChannel(b.bot.Id, user.Id)
	ewrap.Error(err)

	post := &model.Post{
		ChannelId: channel.Id,
		Message:   msg,
	}

	_, _, err = b.Client.CreatePost(post)
	ewrap.Error(err)

}

func (b *Bot) UploadFile(username string, filePath string, filename string) {

	user, _, err := b.Client.GetUserByUsername(username, username)
	ewrap.Error(err)

	channel, _, err := b.Client.CreateDirectChannel(b.bot.Id, user.Id)
	ewrap.Error(err)

	post := &model.Post{
		ChannelId: channel.Id,
	}

	_, _, err = b.Client.UploadFile(core.OpenFiles(filePath), post.ChannelId, filename)
	ewrap.Error(err)

}
