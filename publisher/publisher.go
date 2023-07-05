package publisher

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func publishMessage() {
	// Estabelece a conexão com o servidor RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Falha ao estabelecer conexão com o RabbitMQ: %v", err)
	}
	defer conn.Close()

	// Cria um canal de comunicação
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Falha ao abrir canal: %v", err)
	}
	defer ch.Close()

	// Declara uma fila
	queueName := "minha_fila"
	_, err = ch.QueueDeclare(queueName, false, false, false, false, nil)
	if err != nil {
		log.Fatalf("Falha ao declarar a fila: %v", err)
	}

	// Publica uma mensagem na fila
	message := "Olá, RabbitMQ!"
	err = ch.Publish("", queueName, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	})
	if err != nil {
		log.Fatalf("Falha ao publicar mensagem: %v", err)
	}

	fmt.Println("Mensagem publicada com sucesso!")
}

func main() {
	publishMessage()
}
