package handlers

import (
	"log"
	"mm-bot/internal/mmparser"

	"github.com/gin-gonic/gin"
)

const (
	SLASH_COMMAND = "cqhdt3nnfbrt7bmshrezrd5pge"
)

func hello(c *gin.Context) {

	request, err := mmparser.Parser(c)
	if err != nil {
		log.Println("")
	}

	validate(request.Text, request.UserName)
}
