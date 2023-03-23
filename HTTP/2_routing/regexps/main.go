package main

import (
	"fmt"
	gorillamux "github.com/gorilla/mux"
	"log"
	"net/http"
)

const port = ":9000"

func main() {
	router := gorillamux.NewRouter()
	router.HandleFunc("/prod/{id}", func(writer http.ResponseWriter, request *http.Request) {
		vars := gorillamux.Vars(request)
		fmt.Println(vars["id"])
	})

	// только цифры
	router.HandleFunc("/dig/{id:[0-9]+}", func(writer http.ResponseWriter, request *http.Request) {
		vars := gorillamux.Vars(request)
		fmt.Println(vars["id"])
	})

	mux := http.NewServeMux()
	mux.Handle("/", router)

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal(err)
	}
}
