package main

import (
	"fmt"
	"net/http"

	"short-url/urlshortner"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	urlService := urlshortner.NewService()
	urlshortner.NewHandler(urlService, router)
	fmt.Println("server started")
	http.ListenAndServe(":8082", router)
}
