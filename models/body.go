package models

type WebhookBody struct {
	To      string `json:"to"`
	Content string `json:"content"`
}

type WebhookResponseBody struct {
	Message   string `json:"message"`
	MessageID string `json:"messageId"`
}

type GetSentMessagesResponseBody struct {
	Message  string    `json:"message"`
	Messages []Message `json:"messages"`
}
