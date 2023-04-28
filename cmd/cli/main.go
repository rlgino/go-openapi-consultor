package main

import (
	"encoding/json"
	"fmt"
	"go-openapi-consultor/internal/actions"
	"go-openapi-consultor/internal/client"
	"go-openapi-consultor/internal/domain"
	"os"
)

func main() {
	apiKey := os.Getenv("API_KEY")
	apiOrg := os.Getenv("API_ORG")
	sender := actions.SendMessage{
		Chat: client.NewClient(apiKey, apiOrg),
	}
	msg := domain.Message{Value: "What's the best Argentinian food?"}
	res, err := sender.Chat.Chat(msg)
	if err != nil {
		panic(err)
	}

	resultJSON, _ := json.MarshalIndent(res, "", "  ")
	fmt.Println(string(resultJSON))
}
