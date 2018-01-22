package queue

import (
	"fmt"
	"log"

	"github.com/brunoleonel/email-producer/conf"
	"github.com/streadway/amqp"
)

func getChannel() (channel *amqp.Channel, err error) {

	host := conf.Cfg.Section("").Key("queue_host").Value()
	user := conf.Cfg.Section("").Key("queue_user").Value()
	pass := conf.Cfg.Section("").Key("queue_pass").Value()
	port := conf.Cfg.Section("").Key("queue_port").Value()
	address := fmt.Sprintf("amqp://%s:%s@%s:%s", user, pass, host, port)

	conn, err := amqp.Dial(address)
	defer conn.Close()

	channel, err = conn.Channel()

	if err != nil {
		log.Println("[queue] Houve um erro na conexão com o servidor AMQP: " + err.Error())
		return
	}

	return
}

//AddMessage - Cria uma nova mensagem e adiciona à fila
func AddMessage(json []byte) {

	channel, err := getChannel()
	defer channel.Close()

	queue, err := channel.QueueDeclare(
		"email",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Println("[queue] Houve um erro ao enviar a mensagem para servidor AMQP: " + err.Error())
		return
	}

	err = channel.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        json,
		})

	if err != nil {
		log.Println("[queue] Houve um erro ao enviar a mensagem para servidor AMQP: " + err.Error())
		return
	}

	log.Println("Email enviado para a fila")
	return
}
