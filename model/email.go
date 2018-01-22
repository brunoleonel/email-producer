package model

//Email - representa um e-mail enviado pelo usuário
type Email struct {
	To      string `form:"to"`
	Subject string `form:"subject"`
	Message string `form:"message"`
}
