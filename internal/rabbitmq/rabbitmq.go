package rabbitmq

import (
	"github.com/fvmoraes/api-with-rabbitmq/internal/error"
	"github.com/fvmoraes/api-with-rabbitmq/internal/logs"
	"github.com/streadway/amqp"
)

func ConnectRabbitmq() (*amqp.Connection, *amqp.Channel) {
	// Establish connection to the RabbitMQ server
	conn, err := amqp.Dial("amqp://guest:guest@172.34.0.4:5672/")
	logs.WriteLogFile("INFO", "Connection stabilished")
	error.ValidateError("Failed to establish connection to RabbitMQ", err)

	// Create a communication channel
	ch, err := conn.Channel()
	logs.WriteLogFile("INFO", "Channel opened")
	error.ValidateError("Failed to open channel", err)
	return conn, ch
}
