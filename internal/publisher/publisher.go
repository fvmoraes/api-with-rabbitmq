package publisher

import (
	"math/rand"
	"strconv"

	"github.com/fvmoraes/api-with-rabbitmq/internal/error"
	"github.com/fvmoraes/api-with-rabbitmq/internal/logs"
	"github.com/fvmoraes/api-with-rabbitmq/internal/rabbitmq"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

func PublishMessage() {
	// Establish connection to the RabbitMQ server
	conn, ch := rabbitmq.ConnectRabbitmq()
	defer conn.Close()
	defer ch.Close()

	// Declare an exchange
	exchangeName := "my_exchange"
	err := ch.ExchangeDeclare(exchangeName, "fanout", false, false, false, false, nil)
	error.ValidateError("Failed to declare the exchange:", err)

	// Declare a queue
	queueName := "my_queue"
	_, err = ch.QueueDeclare(queueName, false, false, false, false, nil)
	error.ValidateError("Failed to declare the queue:", err)

	// Bind the exchange to the queue
	err = ch.QueueBind(queueName, "", exchangeName, false, nil)
	error.ValidateError("Failed to bind the queue to the exchange:", err)

	for randomNumber := rand.Intn(10000); randomNumber > 0; randomNumber-- {
		// Publish a message to the exchange
		message := "Rand: " + strconv.Itoa(randomNumber) + ", UUID: " + uuid.New().String()
		err = ch.Publish(exchangeName, "", false, false, amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
		logs.WriteLogFile("INFO", message)
		error.ValidateError("Failed to publish message:", err)
	}
	logs.WriteLogFile("INFO", "All messages published successfully!")
}
