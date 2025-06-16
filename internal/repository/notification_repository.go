package repository

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func SendNotification() (map[string]interface{}, error) {
	var err error
	var client = &http.Client{}
	var data map[string]interface{}

	jsonBody := []byte(`{"client_message": "hello, server!"}`)
	payload := bytes.NewReader(jsonBody)

	request, err := http.NewRequest(http.MethodPost, "http://localhost:3000/notification/send", payload)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
