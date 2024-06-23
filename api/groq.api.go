package api

import (
	"bytes"
	"capstone-project/helper"
	"capstone-project/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func FetchAPI(query string) ([]byte, error) {
	token := helper.GetEnv("GROQ_API_KEY")

	rb := model.GroqRequest{
		Model: "llama3-70b-8192",
		User:  "farismnrr",
		Messages: []struct {
			Role    string `json:"role"`
			Content string `json:"content"`
			Name    string `json:"name"`
		}{
			{
				Role:    "user",
				Content: query,
				Name:    "Faris Munir Mahdi",
			},
		},
	}

	jsonData, err := json.Marshal(rb)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://api.groq.com/openai/v1/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return body, nil
}