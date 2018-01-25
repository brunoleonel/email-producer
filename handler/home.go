package handler

import (
	"net/http"

	"github.com/brunoleonel/email-producer/lib/context"
)

//Home - Retorna a página inicial
func Home(context *context.Context) {
	context.NativeHTML(http.StatusOK, "home")
}
