package main

import (
	"fmt"
	"time"

	"github.com/fvmoraes/api-with-rabbitmq/consumer"
	"github.com/fvmoraes/api-with-rabbitmq/publisher"
)

func main() {
	for true {
		fmt.Println("Working!")

		waitSeconds(10)
		publisher.PublishMessage()

		fmt.Println("Check Rabbit!")

		waitSeconds(10)
		consumer.ConsumeMessages()
	}
}

func waitSeconds(seconds int) {
	fmt.Printf("Wait %d seconds\n", seconds)
	for i := seconds; i > 0; i-- {
		fmt.Println(i, "seconds")
		time.Sleep(1 * time.Second)
	}
}
