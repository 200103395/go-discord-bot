package gpt

import (
	"bot/environment"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
)

/*
	CONNECTION TO GPT CURRENTLY IS NOT DEVELOPED TO WORK WITHOUT PAYMENT TO SERVICE
*/

var Client *resty.Client

const (
	apiEndpoint = "https://api.openai.com/v1/chat/completions"
)

func Start() {
	Client = resty.New()

	log.Println("ChatGPT session started")
}

// Function that handles request and response
func Request(message string) {
	// Creating json request for gpt using ApiKey
	response, err := Client.R().
		SetAuthToken(environment.ApiKey).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"model":      "gpt-3.5-turbo",
			"messages":   []interface{}{map[string]interface{}{"role": "system", "content": message}},
			"max_tokens": 50,
		}).
		Post(apiEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	// Receiving response and decoding from json format
	body := response.Body()
	fmt.Println(string(body))
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error while decoding JSON response:", err)
		return
	}

	// Extract the content from the JSON response
	content := data["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)

	fmt.Println(content)
}
