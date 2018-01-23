package queue

import (
	"fmt"
	"log"

	"github.com/brunoleonel/email-producer/conf"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Printf("%s: %s", msg, err)
		return
	}
}

func getAddress() string {

	host := conf.Cfg.Section("").Key("queue_host").Value()
	user := conf.Cfg.Section("").Key("queue_user").Value()
	pass := conf.Cfg.Section("").Key("queue_pass").Value()
	port := conf.Cfg.Section("").Key("queue_port").Value()
	address := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, pass, host, port)

	return address
}

//AddMessage - Cria uma nova mensagem e adiciona à fila
func AddMessage(json []byte) {

	conn, err := amqp.Dial(getAddress())
	failOnError(err, "[queue] Houve um erro ao conectar com o servidor AMQP")
	defer conn.Close()

	channel, err := conn.Channel()
	failOnError(err, "[queue] Houve um erro na abertura de canal com o servidor AMQP")
	defer channel.Close()

	queue, err := channel.QueueDeclare(
		"email",
		false,
		false,
		false,
		false,
		nil,
	)

	failOnError(err, "[queue] Houve um erro na criação da fila no servidor AMQP")

	err = channel.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        json,
		})

	failOnError(err, "[queue] Houve um erro no envio da mensagem para o servidor AMQP")

	log.Println("Email enviado para a fila")
	return
}
