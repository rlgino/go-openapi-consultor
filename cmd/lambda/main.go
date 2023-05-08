package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"go-openapi-consultor/internal/actions"
	"go-openapi-consultor/internal/client"
	"go-openapi-consultor/internal/domain"
)

func main() {
	apiKey := os.Getenv("API_KEY")
	apiOrg := os.Getenv("API_ORG")
	handler := Handler{
		sender: actions.SendMessage{
			Chat: client.NewClient(apiKey, apiOrg),
		},
	}
	lambda.Start(handler.wrapper)
}

type Handler struct {
	sender actions.SendMessage
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type Response struct {
	Result        []domain.ChatResult `json:"result,omitempty"`
	ErrorResponse *ErrorResponse      `json:"error,omitempty"`
}

func (h Handler) wrapper(event events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	if len(event.Body) == 0 {
		err := fmt.Errorf("bad formatted body")
		return convertResponse(nil, http.StatusBadRequest, err), nil
	}
	messageFrom, err := h.getMessageFrom(event)
	if err != nil {
		return convertResponse(nil, http.StatusBadRequest, err), nil
	}
	response, err := h.execute(messageFrom)
	if err != nil {
		return convertResponse(nil, http.StatusInternalServerError, err), nil
	}
	return convertResponse(response, http.StatusOK, err), nil
}

func (h Handler) getMessageFrom(event events.APIGatewayProxyRequest) (domain.Message, error) {
	msg := domain.Message{}
	err := json.Unmarshal([]byte(event.Body), &msg)
	if err != nil {
		return domain.Message{}, err
	}
	if len(msg.Value) == 0 {
		return domain.Message{}, fmt.Errorf("the message value is required")
	}
	return msg, nil
}

func (h Handler) execute(message domain.Message) ([]domain.ChatResult, error) {
	if len(message.Value) == 0 {
		return nil, fmt.Errorf("the message value is required")
	}
	result, err := h.sender.Chat.Chat(message)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func convertResponse(response []domain.ChatResult, statusCode int, err error) *events.APIGatewayProxyResponse {
	var resp Response
	if err == nil {
		resp.Result = response
	} else {
		resp.ErrorResponse = &ErrorResponse{Message: err.Error()}
	}
	bytes, err := json.Marshal(resp)
	if err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "{\"error\": \"" + err.Error() + "\"}",
		}
	}
	return &events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       string(bytes),
	}
}
