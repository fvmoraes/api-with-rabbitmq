package main

import (
	"fmt"
	"time"

	"github.com/fvmoraes/api-with-rabbitmq/consumer"
	"github.com/fvmoraes/api-with-rabbitmq/publisher"
)

func main() {
	fmt.Println("Working app!")
	callPublishers()
}

func callPublishers() {
	fmt.Println("Call publisher!")
	waitSeconds(5)
	publisher.PublishMessage()

	fmt.Println("Check Rabbit!")
	callConsumers()
}

func callConsumers() {
	fmt.Println("Call consumer!")
	waitSeconds(5)
	consumer.ConsumeMessages()

	fmt.Println("Check Rabbit!")
	callPublishers()
}

func waitSeconds(seconds int) {
	fmt.Printf("Wait %d seconds\n", seconds)
	for i := seconds; i > 0; i-- {
		fmt.Println(i, "seconds")
		time.Sleep(1 * time.Second)
	}
}
