package models

import "time"

type Message struct {
	ID          string    `bson:"_id,omitempty" json:"id"`
	Content     string    `bson:"content,omitempty" json:"content"`
	Number      string    `bson:"number,omitempty" json:"number"`
	IsSent      bool      `bson:"is_sent,omitempty" json:"is_sent"`
	SendingTime time.Time `bson:"sending_time,omitempty" json:"sending_time"`
	MessageID   string    `bson:"message_id,omitempty" json:"message_id"`
}
