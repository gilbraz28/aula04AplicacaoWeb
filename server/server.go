package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gilbraz28/aula04AplicacaoWeb/server/router"
)

func main() {

	router.HandleRoutes()

	http.HandleFunc("/register", register)
	http.HandleFunc("/registerJson", registerJson)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	//http.Handle("/router/pages/", http.StripPrefix("/router/pages/", http.FileServer(http.Dir("router/pages"))))

	http.ListenAndServe(":8080", nil)

}

func register(w http.ResponseWriter, r *http.Request) {

	fmt.Println(" -- DADOS NORMAL FORMULÁRIO -- ")

	fmt.Println(r.PostFormValue("name"))
	fmt.Println(r.FormValue("lastname"))
	fmt.Println(r.FormValue("email"))
	fmt.Println(r.FormValue("password"))
	fmt.Println(r.FormValue("confirm"))
	fmt.Println(r.FormValue("accept"))

	w.WriteHeader(http.StatusOK)
}

func registerJson(w http.ResponseWriter, r *http.Request) {

	type DadosJson struct {
		Name     string `json: "name"`
		Lastname string `json: "lastname"`
		Email    string `json: "email"`
		Password string `json: "password'`
		Confirm  string `json: "confirm"`
		Accept   bool   `json: "accept"`
	}

	var data DadosJson

	json.NewDecoder(r.Body).Decode(&data)

	fmt.Println(" -- Dados Json -- ")

	fmt.Println(data.Name)
	fmt.Println(data.Lastname)
	fmt.Println(data.Email)
	fmt.Println(data.Password)
	fmt.Println(data.Confirm)
	fmt.Println(data.Accept)

	w.WriteHeader(http.StatusOK)
}
