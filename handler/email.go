package handler

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/brunoleonel/email-producer/lib/context"
	"github.com/brunoleonel/email-producer/lib/queue"
	"github.com/brunoleonel/email-producer/model"
)

//ShowForm - Função para exibição da página de envio de e-mail
func ShowForm(context *context.Context) {
	context.Data["to"] = "exemplo@teste.com"
	context.Data["subject"] = "Teste de envio de e-mail"
	context.Data["message"] = "Escreva sua mensagem aqui..."
	context.NativeHTML(200, "form")
	return
}

//SendMail - Função para envio do e-mail
func SendMail(email model.Email, context *context.Context) {

	mails := make([]model.Email, 0)

	admins := []string{"patrao@teste.com", "gerente@teste.com", "subgerente@teste.com"}
	count := 0
	for count < len(admins) {
		mail := model.Email{}
		mail.To = admins[count]
		mail.Subject = fmt.Sprintf("[Monitoramento] Novo e-mail enviado para %s", email.To)
		mail.Message = fmt.Sprintf("Texto enviado: %s", email.Message)
		mails = append(mails, mail)
		count++
	}

	mails = append(mails, email)

	for i := 0; i < len(mails); i++ {
		body, err := json.Marshal(mails[i])
		if err != nil {
			log.Println("[handler/email] Houve um erro: ", err.Error())
		}
		go queue.AddMessage(body)
	}

	context.Redirect("/email")
	return
}
