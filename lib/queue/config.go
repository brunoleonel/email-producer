package queue

import (
	"fmt"
	"log"

	"github.com/brunoleonel/email-producer/conf"
)

//FailOnError - Função para log de erros
func FailOnError(err error, msg string) {
	if err != nil {
		log.Printf("%s: %s", msg, err)
		return
	}
}

//GetAddress - Função para recuperar o endereço do servidor AMQP
func GetAddress() string {

	host := conf.Cfg.Section("").Key("queue_host").Value()
	user := conf.Cfg.Section("").Key("queue_user").Value()
	pass := conf.Cfg.Section("").Key("queue_pass").Value()
	port := conf.Cfg.Section("").Key("queue_port").Value()
	address := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, pass, host, port)

	return address
}
