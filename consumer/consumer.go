package consumer

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func ConsumeMessages() {
	// Establish connection to the RabbitMQ server
	conn, err := amqp.Dial("amqp://guest:guest@172.34.0.4:5672/")
	if err != nil {
		log.Fatalf("Failed to establish connection to RabbitMQ: %v", err)
	}
	defer conn.Close()

	// Create a communication channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open channel: %v", err)
	}
	defer ch.Close()

	// Declare a queue
	queueName := "my_queue"
	_, err = ch.QueueDeclare(queueName, false, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare the queue: %v", err)
	}

	// Consume messages from the queue
	msgs, err := ch.Consume(queueName, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to consume messages: %v", err)
	}

	// Process received messages
	for msg := range msgs {
		fmt.Printf("Received message: %s\n", msg.Body)
	}
}
