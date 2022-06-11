package main

import (
	"net/http"

	"github.com/gilbraz28/aula04AplicacaoWeb/server/src/login"
)

func main() {

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("pages"))))

	http.ListenAndServe(":8080", nil)

	login.Login("t")

}
