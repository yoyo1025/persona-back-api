package database

import(
	"encoding/json"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"net/http"
)


var openaiClient *openai.Client

func SetOpenAIClient(client *openai.Client) {
	openaiClient = client
}



type Persona struct{
	Name string `json:"name"`
	user_id int64 `"json:user_id"`
	Sex	string `"json:sex"`
	Age	int64 `"json:age"`
	Profession string `"json:profession"`
	Problems string `"json:problems"`
	Behavior string `"json:behavior"`
	

}

//通信処理
func InputPersona(w http.ResponseWriter, r *http.Request){
	//postメソッド以外をはじく
	if r.Method != http.MethodPost {
		http.Error(w, "許可されていないメソッドです", http.StatusMethodNotAllowed)
		return
	}
	fmt.Println("InputPersonaが呼ばれました")
	
	//ペルソナ型の構造体を使用
	var persona Persona
// リクエストボディからJSONデータをデコード
	err := json.NewDecoder(r.Body).Decode(&persona)
	if err != nil {
		http.Error(w, "リクエストデータのデコードに失敗しました", http.StatusBadRequest)
		return
	}
	//登録画面が空白の場合の送信拒否
	if persona.Name == "" || persona.Sex == "" || persona.Problems =="" || persona.Behavior == "" || persona.Profession == "" {
		http.Error(w, "必要なフィールドが不足しています", http.StatusBadRequest)
		return
	}

		// データベースに挿入
	query := `  INSERT INTO persona (name, user_id, sex, age, profession, problems, behavior) VALUES ($1, $2, $3, $4, $5, $6, $7) returning id`
	var insertedID int
	err = db.QueryRow(query, persona.Name, 1, persona.Sex, persona.Age, persona.Profession, persona.Problems, persona.Behavior).Scan(&insertedID)
	if err != nil {
			http.Error(w, "データベースへの挿入に失敗しました: "+err.Error(), http.StatusInternalServerError)
			return
	}
	fmt.Println("データベースの挿入に成功しました")
	fmt.Fprintln(w, "登録が完了しました！")
}

