package consumer

import (
	"fmt"

	"github.com/fvmoraes/api-with-rabbitmq/internal/error"
	"github.com/fvmoraes/api-with-rabbitmq/internal/logs"
	"github.com/fvmoraes/api-with-rabbitmq/internal/rabbitmq"
)

func ConsumeMessages() {
	fmt.Println("Consuming messages")
	conn, ch := rabbitmq.ConnectRabbitmq()
	defer conn.Close()
	defer ch.Close()

	// Declare a queue
	queueName := "my_queue"
	_, err := ch.QueueDeclare(queueName, false, false, false, false, nil)
	error.ValidateError("Failed to declare the queue", err)

	// Consume messages from the queue
	msgs, err := ch.Consume(queueName, "", true, false, false, false, nil)

	error.ValidateError("Failed to consume messages", err)

	// Process received messages
	for msg := range msgs {
		logs.WriteLogFile("INFO", "Channel opened")
		logs.WriteLogFile("INFO", "Received message: "+string(msg.Body))
	}
}
