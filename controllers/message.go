package controllers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mertcankirtay/message-service/db"
	"github.com/mertcankirtay/message-service/models"
	"github.com/mertcankirtay/message-service/services"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// getSentMessages godoc
// @Summary Get sent messages
// @Description Get sent messages
// @Accept  json
// @Produce  json
// @Query page
// @Router /messages [get]
func GetSentMessages(c *gin.Context) {
	// Get desired page from querystring
	pageQuery := c.Query("page")
	page, err := strconv.Atoi(pageQuery)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Page query should be a number"})
		return
	}

	// Set options for pagination
	opt := options.Find().SetLimit(10)
	opt.SetSkip(int64(page) * 10)

	// Fetch messages from db
	cur, err := db.MessageColl.Find(context.TODO(), bson.M{"is_sent": true})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error!"})
	}

	// Prepare response body
	respBody := models.GetSentMessagesResponseBody{
		Message:  "OK",
		Messages: make([]models.Message, 0),
	}

	for cur.Next(context.TODO()) {
		var msg models.Message
		if err = cur.Decode(&msg); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error!"})
			return
		}

		respBody.Messages = append(respBody.Messages, msg)
	}

	// Send response
	c.AbortWithStatusJSON(http.StatusOK, respBody)
}

// toggleScheduler godoc
// @Summary Toggle the scheduler
// @Description Toggle the scheduler
// @Accept  json
// @Produce  json
// @Router /scheduler/toggle [post]
func ToggleScheduler(c *gin.Context) {
	if services.IsRunning {
		services.StopScheduler()
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "OK"})
		return
	}

	services.InitScheduler()
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "OK"})
}
