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

//通信処理
func InputPersona(w http.ResponseWriter, r *http.Request){
	//getメソッド以外をはじく
	if r.Method != http.MethodGet {
		http.Error(w, "許可されていないメソッドです", http.StatusMethodNotAllowed)
		return
	}
	
	//ペルソナ型の構造体を使用
	var persona Persona
// リクエストボディからJSONデータをデコード
	err := json.NewDecoder(r.Body).Decode(&persona)
	if err != nil {
		http.Error(w, "リクエストデータのデコードに失敗しました", http.StatusBadRequest)
		return
	}
	if persona.Name == "" || persona.Sex == "" || persona.Age == "" || persona.Problems =="" || persona.Behavior == "" || persona.Profession == "" {
		http.Error(w, "必要なフィールドが不足しています", http.StatusBadRequest)
		return
	}

}

