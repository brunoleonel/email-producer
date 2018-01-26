package queue

import (
	"log"

	"github.com/streadway/amqp"
)

//AddConcurrentMessage - Cria uma nova mensagem e adiciona à fila
func AddConcurrentMessage(json []byte) {

	conn, err := amqp.Dial(GetAddress())
	FailOnError(err, "[queue] Houve um erro ao conectar com o servidor AMQP")
	defer conn.Close()

	channel, err := conn.Channel()
	FailOnError(err, "[queue] Houve um erro na abertura de canal com o servidor AMQP")
	defer channel.Close()

	queue, err := channel.QueueDeclare(
		"email",
		false,
		false,
		false,
		false,
		nil,
	)

	FailOnError(err, "[queue] Houve um erro na criação da fila no servidor AMQP")

	err = channel.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         json,
		})

	FailOnError(err, "[queue] Houve um erro no envio da mensagem para o servidor AMQP")

	log.Println("Email enviado para a fila")
	return
}
