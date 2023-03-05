package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	port       = ":9000"
	queryParam = "key"

	adminUsername = "admin"
	adminPassword = "admin"
)

//fmt.Println("HEADERS: ", request.Header)
//fmt.Println("URL: ", request.URL)
//fmt.Println("URL Query: ", request.URL.Query())

func main() {
	implementation := server{
		data: map[string]string{},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", authMiddleware(rootHandler(implementation)))

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal(err)
	}
}

func authMiddleware(handler http.Handler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		username, password, ok := request.BasicAuth()
		if !ok {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}

		if username != adminUsername || password != adminPassword {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}

		handler.ServeHTTP(writer, request)
	}
}

func rootHandler(implementation server) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			implementation.Read(writer, request)

		case http.MethodPost:
			implementation.Create(writer, request)

		case http.MethodPut:
			implementation.Update(writer, request)

		case http.MethodDelete:
			implementation.Delete(writer, request)

		default:
			fmt.Printf("unsupported method: [%s]", request.Method)
		}
	}
}

// not thread safe
type server struct {
	data map[string]string
}

func (s *server) Create(writer http.ResponseWriter, request *http.Request) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	type data struct {
		Key   string
		Value string
	}

	var unmarshalled data
	if err := json.Unmarshal(body, &unmarshalled); err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	if unmarshalled.Key == "" || unmarshalled.Value == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if _, ok := s.data[unmarshalled.Key]; ok {
		writer.WriteHeader(http.StatusConflict)
		return
	}

	s.data[unmarshalled.Key] = unmarshalled.Value
}

func (s *server) Read(writer http.ResponseWriter, request *http.Request) {
	key := request.URL.Query().Get(queryParam)
	if key == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	value, ok := s.data[key]
	if !ok {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	if _, err := writer.Write([]byte(value)); err != nil {
		log.Printf("error while writing body, err: [%s]", err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	//writer.WriteHeader(http.StatusOK) // default
}

func (s *server) Update(writer http.ResponseWriter, request *http.Request) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	type data struct {
		Key   string
		Value string
	}

	var unmarshalled data
	if err := json.Unmarshal(body, &unmarshalled); err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	if unmarshalled.Key == "" || unmarshalled.Value == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if _, ok := s.data[unmarshalled.Key]; !ok {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	s.data[unmarshalled.Key] = unmarshalled.Value
}

func (s *server) Delete(writer http.ResponseWriter, request *http.Request) {
	key := request.URL.Query().Get(queryParam)
	if key == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if _, ok := s.data[key]; !ok {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	delete(s.data, key)
}
