package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/mertcankirtay/message-service/controllers"
	"github.com/mertcankirtay/message-service/db"
	"github.com/mertcankirtay/message-service/helpers"
	"github.com/mertcankirtay/message-service/services"
)

func main() {
	// Init everything
	errCh := make(chan error, 2)

	go func() { errCh <- db.InitMongo() }()
	go func() { errCh <- db.InitRedis() }()

	// Check if database connecitons are successful
	for range 2 {
		if err := <-errCh; err != nil {
			fmt.Printf("UNABLE TO INIT SERVICE;\n%s\n", err.Error())
			exit()
			return
		}
	}

	close(errCh)

	// Fetch vars from env
	if err := helpers.InitVars(); err != nil {
		fmt.Printf("UNABLE TO INIT SERVICE;\n%s\n", err.Error())
		exit()
		return
	}

	// Init message scheduler
	go services.InitScheduler()

	// Listen exit signals
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	go func() { <-sigCh; exit(); os.Exit(1) }()

	// Finally init api router
	controllers.InitRouting()
}

func exit() {
	fmt.Println("Exiting...")
	db.DisconnectMongo()
	db.DisconnectRedis()
}
