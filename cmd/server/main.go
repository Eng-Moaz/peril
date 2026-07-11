package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Starting Peril server...")
	const connectionURL = "amqp://guest:guest@localhost:5672/"
	conn, err := amqp.Dial(connectionURL)
	if err != nil {
		log.Fatal("Error in connecting: ", err)
	}
	defer conn.Close()
	fmt.Println("Connection started successfully..")

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
	fmt.Println("Program is shutting down")
}
