package controllers

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Message Service API
// @version 1.0
// @description This service sends automated messages every 2 minutes.
// @host localhost:8000
// @BasePath /api/v1

func InitRouting() {
	app := gin.Default()

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	app.GET("/api/v1/messages", GetSentMessages)
	app.POST("/api/v1/scheduler/toggle", ToggleScheduler)

	app.Run(fmt.Sprintf(":%s", os.Getenv("SERVICE_PORT")))
}
