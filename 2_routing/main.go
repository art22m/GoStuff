package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":9000"

func main() {
	mux := http.NewServeMux()

	// многоуровневый
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Hello 9000")
	})

	// одноуровненый
	mux.HandleFunc("/article", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Article single level")
	})

	// многоуровневый
	mux.HandleFunc("/article/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Article multilevel")
	})

	mux.Handle("/home", http.RedirectHandler("http://localhost:9000/article", http.StatusPermanentRedirect))

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal(err)
	}
}
