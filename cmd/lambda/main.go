package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"go-openapi-consultor/internal/actions"
	"go-openapi-consultor/internal/client"
	"go-openapi-consultor/internal/domain"
	"os"
)

func main() {
	apiKey := os.Getenv("API_KEY")
	apiOrg := os.Getenv("API_ORG")
	handler := Handler{
		sender: actions.SendMessage{
			Chat: client.NewClient(apiKey, apiOrg),
		},
	}

	lambda.Start(handler.HandleRequest)
}

type Handler struct {
	sender actions.SendMessage
}

func (h Handler) HandleRequest(_ context.Context, message domain.Message) ([]domain.ChatResult, error) {
	result, err := h.sender.Chat.Chat(message)
	if err != nil {
		return nil, err
	}
	return result, nil
}
