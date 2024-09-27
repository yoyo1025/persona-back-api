package database

import "github.com/sashabaranov/go-openai"

var openaiClient *openai.Client

func SetOpenAIClient(client *openai.Client) {
	openaiClient = client
}