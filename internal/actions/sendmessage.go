package actions

import (
	"go-openapi-consultor/internal/domain"
)

type SendMessage struct {
	Chat domain.SmartChat
}

func (sender SendMessage) send(message domain.Message) ([]domain.ChatResult, error) {
	result, err := sender.Chat.Chat(message)
	if err != nil {
		return nil, err
	}
	return result, nil
}
