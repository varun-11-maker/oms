package main

import (
	"encoding/json"
	"fmt"
	"log"
	"usingPostgres/middleware"
	"usingPostgres/models"

	"github.com/streadway/amqp"
)

// getting all of the orders
var stocks []models.Stock = middleware.Stocke

func MQConnect() (*amqp.Connection, *amqp.Channel, error) {
	// Connect to RabbitMQ
	url := "amqp://guest:guest@localhost:5672/"
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, nil, err
	}
	channel, err := conn.Channel()
	if err != nil {
		return nil, nil, err
	}
	_, err = channel.QueueDeclare(
		"email_queue",
		true,  // Durable (survive restarts)
		false, // Delete when unused
		false, // Exclusive (only this connection can use it)
		false, // No-wait
		nil,   // Arguments
	)
	if err != nil {
		return nil, nil, err
	}

	return conn, channel, nil
}
func MQPublish(channel *amqp.Channel, message []byte) error {
	err := channel.Publish(
		"",
		"email_queue",
		false,
		false, // Immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	)
	return err
}
func main() {
	conn, channel, err := MQConnect()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	defer channel.Close()
	fmt.Println(stocks)
	message, err := json.Marshal(stocks)
	if err != nil {
		log.Fatal(err)
	}
	err = MQPublish(channel, message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Message sent:", string(message))
}
