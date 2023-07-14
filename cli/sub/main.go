package main

import (
	"fmt"
	"time"

	"github.com/fvmoraes/api-with-rabbitmq/internal/consumer"
)

func main() {
	fmt.Println("Working sub!")
	waitSeconds(5)
	consumer.ConsumeMessages()
	fmt.Println("Check Rabbit!")
}

func waitSeconds(seconds int) {
	fmt.Printf("Wait %d seconds\n", seconds)
	for i := seconds; i > 0; i-- {
		fmt.Println(i, "seconds")
		time.Sleep(1 * time.Second)
	}
}
