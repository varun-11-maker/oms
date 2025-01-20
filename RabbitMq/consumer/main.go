package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// Define the message structure
type Stock struct {
	StockID int64  `json:"stockid"`
	Name    string `json:"name"`
	Price   int64  `json:"price"`
	Company string `json:"company"`
}

// Function to connect to RabbitMQ
func MQConnect() (*amqp.Connection, *amqp.Channel, error) {
	// Connect to RabbitMQ
	url := "amqp://guest:guest@localhost:5672/"
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, nil, err
	}

	// Create a channel
	channel, err := conn.Channel()
	if err != nil {
		return nil, nil, err
	}

	_, err = channel.QueueDeclare(
		"email_queue", // Queue name
		true,          // Durable
		false,         // Delete when unused
		false,         // Exclusive
		false,         // No-wait
		nil,           // Arguments
	)
	if err != nil {
		return nil, nil, err
	}

	return conn, channel, nil
}

func MQConsume(channel *amqp.Channel) error {
	msgs, err := channel.Consume(
		"email_queue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	for msg := range msgs {
		var stock Stock
		err := json.Unmarshal(msg.Body, &stock)
		if err != nil {
			log.Println("Error unmarshalling message:", err)
			continue
		}
		fmt.Printf("Id: %d\n", stock.StockID)
		fmt.Printf("Message: %s\n", stock.Name)
		fmt.Printf("Price: %d\n", stock.Price)
		fmt.Printf("Company: %s\n", stock.Company)
	}

	return nil
}

func main() {
	// Connect to RabbitMQ
	conn, channel, err := MQConnect()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	defer channel.Close()

	// Start consuming messages
	err = MQConsume(channel)
	if err != nil {
		log.Fatal(err)
	}
}
