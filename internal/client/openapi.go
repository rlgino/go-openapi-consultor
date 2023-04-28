package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-openapi-consultor/internal/domain"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type client struct {
	ApiKey       string
	Organization string
}

func (c *client) Chat(message domain.Message) ([]domain.ChatResult, error) {
	r := CreateCompletionsRequest{
		Model: "gpt-3.5-turbo",
		Messages: []Message{
			{
				Role:    "user",
				Content: message.Value,
			},
		},
		Temperature: 0.7,
	}

	completions, err := c.CreateCompletions(r)
	if err != nil {
		return nil, err
	}
	if completions.Error != nil {
		return nil, fmt.Errorf("[ERROR][CODE: %s][%s]", completions.Error.Code, completions.Error.Message)
	}

	var result []domain.ChatResult
	for _, choice := range completions.Choices {
		for _, item := range strings.Split(choice.Message.Content, "\n") {
			if len(strings.TrimSpace(item)) > 0 {
				result = append(result, domain.ChatResult{
					Message: item,
				})
			}
		}
	}
	return result, nil
}

// NewClient creates a new client
func NewClient(apiKey string, organization string) domain.SmartChat {
	return &client{
		ApiKey:       apiKey,
		Organization: organization,
	}
}

// Post makes a post request
func (c *client) Post(url string, input any) (response []byte, err error) {
	response = make([]byte, 0)

	rJson, err := json.Marshal(input)
	if err != nil {
		return response, err
	}

	resp, err := c.Call(http.MethodPost, url, bytes.NewReader(rJson))
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	response, err = io.ReadAll(resp.Body)
	return response, err
}

// Get makes a get request
func (c *client) Get(uri string, input ...string) (response []byte, err error) {
	if input != nil {
		values := strings.Join(input, "&")
		query := url.QueryEscape(values)

		if query != "" {
			uri += "?" + query
		}
	}

	resp, err := c.Call(http.MethodGet, uri, nil)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	response, err = io.ReadAll(resp.Body)
	return response, err
}

// Call makes a request
func (c *client) Call(method string, url string, body io.Reader) (response *http.Response, err error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return response, err
	}

	req.Header.Add("Authorization", "Bearer "+c.ApiKey)
	req.Header.Add("Content-Type", "application/json")
	if c.Organization != "" {
		req.Header.Add("OpenAI-Organization", c.Organization)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	return resp, err
}
