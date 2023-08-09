package publisher

import (
	"encoding/json"

	"github.com/fvmoraes/api-with-rabbitmq/internal/error"
	"github.com/fvmoraes/api-with-rabbitmq/internal/logs"
	"github.com/fvmoraes/api-with-rabbitmq/internal/models"
	"github.com/fvmoraes/api-with-rabbitmq/internal/rabbitmq"
	"github.com/streadway/amqp"
)

func PublishMessage(foobar *models.Foobar, action string) {
	// Establish connection to the RabbitMQ server
	conn, ch := rabbitmq.ConnectRabbitmq()
	defer conn.Close()
	defer ch.Close()

	// Declare an exchange
	exchangeName := action + "_exchange"
	err := ch.ExchangeDeclare(exchangeName, "fanout", true, false, false, false, nil)
	error.ValidateError("Failed to declare the exchange:", err)

	// Declare a queue
	queueName := action + "_queue"
	_, err = ch.QueueDeclare(queueName, true, false, false, false, nil)
	error.ValidateError("Failed to declare the queue:", err)

	// Bind the exchange to the queue
	err = ch.QueueBind(queueName, "", exchangeName, false, nil)
	error.ValidateError("Failed to bind the queue to the exchange:", err)

	// Serialize the foobar object to JSON
	messageBytes, err := json.Marshal(foobar)
	if err != nil {
		error.ValidateError("Failed to serialize foobar to JSON:", err)
		return
	}

	// Publish the serialized message to the exchange
	err = ch.Publish(exchangeName, "", false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        messageBytes,
	})
	error.ValidateError("Failed to publish message:", err)
	logs.WriteLogFile("INFO", "Message published successfully!")
}
