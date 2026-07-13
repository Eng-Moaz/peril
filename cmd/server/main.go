package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bootdotdev/learn-pub-sub-starter/internal/pubsub"
	"github.com/bootdotdev/learn-pub-sub-starter/internal/routing"
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

	chann, err := conn.Channel()
	if err != nil {
		log.Fatal("Error in creating channel: ", err)
	}

	exchange := routing.ExchangePerilDirect
	key := routing.PauseKey

	dataToSend := routing.PlayingState{
		IsPaused: true,
	}

	err = pubsub.PublishJSON(chann, exchange, key, dataToSend)
	if err != nil {
		log.Fatal("Error in publishing: ", err)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
	fmt.Println("Program is shutting down")
}
