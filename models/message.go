package models

import "time"

type Message struct {
	ID          string    `bson:"_id,omitempty"`
	Content     string    `bson:"content,omitempty"`
	Number      string    `bson:"number,omitempty"`
	IsSent      bool      `bson:"is_sent,omitempty"`
	SendingTime time.Time `bson:"sending_time,omitempty"`
	MessageID   string    `bson:"message_id,omitempty"`
}
