package controllers

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func InitRouting() {
	app := gin.Default()

	app.GET("/api/v1/messages", GetSentMessages)
	app.POST("/api/v1/scheduler/toggle", ToggleScheduler)

	app.Run(fmt.Sprintf(":%s", os.Getenv("SERVICE_PORT")))
}
