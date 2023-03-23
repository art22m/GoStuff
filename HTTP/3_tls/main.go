package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":9000"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("mux handle with tls")
	})

	if err := http.ListenAndServeTLS(port, "./server.crt", "./server.key", mux); err != nil {
		log.Fatal(err)
	}
}
