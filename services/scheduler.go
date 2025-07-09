package services

import (
	"context"
	"time"
)

var cancel context.CancelFunc

func InitScheduler() {
	var ctx context.Context
	ctx, cancel = context.WithCancel(context.Background())

outer:
	for {
		// Check if the context is cancelled
		select {
		case <-ctx.Done():
			break outer
		default:
			sendMessages()
			time.Sleep(time.Minute * 2)
		}

	}
}

func StopScheduler() {
	cancel()
}
