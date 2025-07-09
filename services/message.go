package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/mertcankirtay/message-service/db"
	"github.com/mertcankirtay/message-service/helpers"
	"github.com/mertcankirtay/message-service/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func sendMessages() {
	// Get all unsend messages
	cur, err := db.MessageColl.Find(context.TODO(), bson.M{"is_sent": false}, options.Find().SetLimit(2))

	if err != nil {
		fmt.Printf("Error while fetching messages;\n%s\n", err.Error())
		return
	}

	// Decoder message records and trigger send
	for cur.Next(context.TODO()) {
		var msg models.Message
		if err := cur.Decode(&msg); err != nil {
			fmt.Printf("Error while decoding message record;\n%s\n", err.Error())
			continue
		}

		// Check content length
		if len(msg.Content) > 140 {
			fmt.Printf("Error while sending message \"Content is too long\";\n%s\n", err.Error())
			continue
		}

		go sendMessage(msg.ID, msg.Number, msg.Content)
	}

}

func sendMessage(id, phone, content string) {
	// Prepare body
	body := models.WebhookBody{
		To:      phone,
		Content: content,
	}

	// Encode body value to json
	encodedBody, err := json.Marshal(body)

	if err != nil {
		fmt.Printf("Could not send message;\n%s\n", err.Error())
		return
	}

	// Prepare json for http request
	bodyReader := bytes.NewBuffer(encodedBody)

	// Create request to send
	req, err := http.NewRequest(http.MethodPost, helpers.URL, bodyReader)

	if err != nil {
		fmt.Printf("Could not send message;\n%s\n", err.Error())
		return
	}

	// Set headers for request
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-ins-auth-key", helpers.AuthKey)

	// Send request
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Printf("Could not send message;\n%s\n", err.Error())
		return
	}

	defer resp.Body.Close()

	// Read response body
	respBody, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Could not send message;\n%s\n", err.Error())
		return
	}

	var response models.WebhookResponseBody

	// Decode response body
	if err = json.Unmarshal(respBody, &response); err != nil {
		fmt.Printf("Could not send message;\n%s\n", err.Error())
		return
	}

	if resp.StatusCode != http.StatusAccepted {
		fmt.Printf("Error from webhook;\n%s\n", response.Message)
		return
	}

	// Save result to db
	saveResult(id, response.MessageID)
}

func saveResult(id, msgID string) {
	// Create document to save
	newDoc := models.Message{
		IsSent:      true,
		MessageID:   msgID,
		SendingTime: time.Now(),
	}

	// Update db
	if _, err := db.MessageColl.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": newDoc}); err != nil {
		fmt.Printf("Error while saving message result;\n%s\n", err.Error())
		return
	}

	// Prepare doc for redis

	newDoc.ID = id
	encodedDoc, err := bson.Marshal(newDoc)

	if err != nil {
		fmt.Printf("Error while saving message result;\n%s\n", err.Error())
		return
	}

	// Add sent message to redis cache
	if _, err = db.RedisClient.Set(context.TODO(), msgID, encodedDoc, 0).Result(); err != nil {
		fmt.Printf("Error while saving message result;\n%s\n", err.Error())
		return
	}
}
