package consumer

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func consumeMessages() {
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

	// Consome mensagens da fila
	msgs, err := ch.Consume(queueName, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Falha ao consumir mensagens: %v", err)
	}

	// Processa as mensagens recebidas
	for msg := range msgs {
		fmt.Printf("Mensagem recebida: %s\n", msg.Body)
	}
}

func main() {
	consumeMessages()
}
