package util

import (
	"log"
	"os"

	"github.com/sashabaranov/go-openai"
)


func InitOpenAI(client **openai.Client) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OpenAI APIキーが設定されていません")
	}
	*client = openai.NewClient(apiKey)
}
