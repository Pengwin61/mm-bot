package mmparser

import (
	"fmt"
	"net/url"

	"github.com/gin-gonic/gin"
)

type Request struct {
	ChannelID   string `json:"channel_id"`
	ChannelName string `json:"channel_name"`
	ResponseURL string `json:"response_url"`
	Text        string `json:"text"`
	Token       string `json:"token"`
	TeamID      string `json:"team_id"`
	TeamDomain  string `json:"team_domain"`
	TriggerID   string `json:"trigger_id"`
	UserID      string `json:"user_id"`
	UserName    string `json:"user_name"`
	Command     string `json:"command"`
}

func Parser(c *gin.Context) (*Request, error) {
	body, err := c.GetRawData()
	if err != nil {
		c.JSON(500, gin.H{"error": "Произошла ошибка"})
		return nil, err
	}

	values, err := url.ParseQuery(string(body))
	if err != nil {
		fmt.Println("Failed to parse query:", err)
		return nil, err
	}
	command := Request{
		ChannelID:   values.Get("channel_id"),
		ChannelName: values.Get("channel_name"),
		Command:     values.Get("command"),
		ResponseURL: values.Get("response_url"),
		TeamDomain:  values.Get("team_domain"),
		TeamID:      values.Get("team_id"),
		Text:        values.Get("text"),
		Token:       values.Get("token"),
		TriggerID:   values.Get("trigger_id"),
		UserID:      values.Get("user_id"),
		UserName:    values.Get("user_name"),
	}
	return &command, nil

}
