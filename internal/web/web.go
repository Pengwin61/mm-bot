package web

import (
	"fmt"
	"mm-bot/internal/web/handlers"

	"github.com/gin-gonic/gin"
)

func InitGin() {
	r := gin.Default()
	handlers.InitHandlers(r)

	// RUN
	err := r.Run(fmt.Sprint("0.0.0.0", ":", "8080"))
	if err != nil {
		panic(err)
	}
}
