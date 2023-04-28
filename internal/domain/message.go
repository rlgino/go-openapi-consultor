package domain

type Message struct {
	Value string
}

type ChatResult struct {
	Message string
}

type SmartChat interface {
	Chat(Message) ([]ChatResult, error)
}
