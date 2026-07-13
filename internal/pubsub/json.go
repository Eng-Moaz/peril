package pubsub

import (
	"context"
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func PublishJSON[T any](ch *amqp.Channel, exchange, key string, val T) error{
	data, err := json.Marshal(val)
	if err != nil{
		return fmt.Errorf("Error in marshalling %v", err)
	}

	msg := amqp.Publishing{
		ContentType: "application/json",
		Body: data,
	}
	err = ch.PublishWithContext(context.Background(), exchange, key, false, false, msg)
	if err != nil{
		return fmt.Errorf("Error in publishing with context %v", err)
	}
	return nil	
}
