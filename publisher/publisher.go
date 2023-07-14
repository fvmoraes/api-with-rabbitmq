package publisher

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"

	"github.com/streadway/amqp"
)

func PublishMessage() {
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

	// Declare an exchange
	exchangeName := "my_exchange"
	err = ch.ExchangeDeclare(exchangeName, "fanout", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare the exchange: %v", err)
	}

	// Declare a queue
	queueName := "my_queue"
	_, err = ch.QueueDeclare(queueName, false, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare the queue: %v", err)
	}

	// Bind the exchange to the queue
	err = ch.QueueBind(queueName, "", exchangeName, false, nil)
	if err != nil {
		log.Fatalf("Failed to bind the queue to the exchange: %v", err)
	}

	for randomNumber := rand.Intn(10000); randomNumber > 0; randomNumber-- {
		fmt.Println("Publishing " + strconv.Itoa(randomNumber) + " messages")

		// Publish a message to the exchange
		message := "Hello, RabbitMQ! message " + strconv.Itoa(randomNumber)
		err = ch.Publish(exchangeName, "", false, false, amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
		if err != nil {
			log.Fatalf("Failed to publish message: %v", err)
		}
	}

	fmt.Println("Messages published successfully!")
}
