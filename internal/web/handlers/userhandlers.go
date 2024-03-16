package handlers

import (
	"fmt"
	"log"
	"mm-bot/internal/mmparser"

	"github.com/gin-gonic/gin"
)

const (
	SLASH_COMMAND = "cqhdt3nnfbrt7bmshrezrd5pge"
)

func hello(c *gin.Context) {

	requerst, err := mmparser.Parser(c)
	if err != nil {
		log.Println("")
	}

	fmt.Println("ChannelID", requerst.ChannelID)
	fmt.Println("ChannelName", requerst.ChannelName)
	fmt.Println("Command", requerst.Command)
	fmt.Println("ResponseURL", requerst.ResponseURL)
	fmt.Println("TeamDomain", requerst.TeamDomain)
	fmt.Println("TeamID", requerst.TeamID)
	fmt.Println("Text", requerst.Text)
	fmt.Println("Token", requerst.Token)
	fmt.Println("ChannelName", requerst.ChannelName)
	fmt.Println("ChannelName", requerst.ChannelName)

}
