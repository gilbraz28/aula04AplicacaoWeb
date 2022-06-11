package main

import (
	"net/http"
)

func main() {

	http.Handle("/router/assets/", http.StripPrefix("/router/assets/", http.FileServer(http.Dir("router/assets"))))

	http.Handle("/router/pages/", http.StripPrefix("/router/pages/", http.FileServer(http.Dir("router/pages"))))

	http.ListenAndServe(":8080", nil)

	//router.HandleRoutes()

}
