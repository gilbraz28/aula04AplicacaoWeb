package router

import (
	"net/http"

	"github.com/gilbraz28/aula04AplicacaoWeb/server/router/account"
	"github.com/gilbraz28/aula04AplicacaoWeb/server/router/login"
)

func HandleRoutes(w http.ResponseWriter, r *http.Request) {

	http.HandleFunc("/", login.Index)

	http.HandleFunc("/", account.Index)

}
