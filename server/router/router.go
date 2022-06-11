package router

import (
	"net/http"

	"github.com/gilbraz28/aula04AplicacaoWeb/server/router/account"
	"github.com/gilbraz28/aula04AplicacaoWeb/server/router/login"
)

func HandleRoutes() {

	http.HandleFunc("/", login.Index)

	http.HandleFunc("/account", account.Index)

}
