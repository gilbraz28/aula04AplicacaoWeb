package main

import (
	"net/http"
)

func main() {

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("pages"))))

	http.ListenAndServe(":8080", nil)

}
