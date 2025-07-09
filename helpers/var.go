package helpers

import (
	"errors"
	"os"
)

var URL, AuthKey string

func InitVars() (err error) {
	URL = os.Getenv("WEBHOOK_URL")
	AuthKey = os.Getenv("AUTH_KEY")

	if URL == "" || AuthKey == "" {
		err = errors.New("WEBHOOK_URL and AUTH_KEY is required")
	}

	return
}
