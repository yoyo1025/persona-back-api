package database

import "github.com/sashabaranov/go-openai"

var openaiClient *openai.Client

func SetOpenAIClient(client *openai.Client) {
	openaiClient = client
}

import(
	"encoding/json"
	"fmt"
)

type Persona struct{
	Name string `json:"name"`
	Sex	string `"json:sex"`
	Age	int `"json:age"`
	Profession string `"json:profession"`
	Problems string `"json:problems"`
	Behavior string `"json:behavior"`
}



