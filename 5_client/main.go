package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}

	request, err := http.NewRequest(http.MethodGet, "http://localhost:9000?key=client_test", nil)
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Add("Authorization", "Basic YWRtaW46YWRtaW4=")

	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}

func create() {
	client := &http.Client{}

	type data struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}

	reqData := data{
		Key:   "client_test",
		Value: "client_pass",
	}

	reqBody, err := json.Marshal(&reqData)
	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest(http.MethodPost, "http://localhost:9000", bytes.NewReader(reqBody))
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Add("Authorization", "Basic YWRtaW46YWRtaW4=")

	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.StatusCode)
}
