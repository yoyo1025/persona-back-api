package database


import (
"github.com/sashabaranov/go-openai"
"fmt"
"net/http"
"encoding/json"
"github.com/yoyo1025/persona-back-api/model"
"github.com/yoyo1025/persona-back-api/util"
) 


var openaiClient *openai.Client

func SetOpenAIClient(client *openai.Client) {
	openaiClient = client
}


type User struct {
    Id        int      `json:"id"`
    User_id   int      `json:"user_id"`
    Name      string   `json:"name"`
	Problems  string   `json:"problem"`
}


func UseridHandler(w http.ResponseWriter, r *http.Request){

	if r.Method != http.MethodGet {
		http.Error(w, "許可されていないメソッドです", http.StatusMethodNotAllowed)
		return
	}
	// クエリ実行
	query := `SELECT id,user_id,name,problems FROM persona WHERE user_id = $1`
	rows, err := db.Query(query, 1)
	if err != nil {
  	http.Error(w, "コメントの取得に失敗しました: "+err.Error(), http.StatusInternalServerError)
  		return
	}

	defer rows.Close()


	
	u := User{}
	var result []User

	for rows.Next() {
		


		if err := rows.Scan(&u.Id, &u.User_id, &u.Name, &u.Problems); err != nil {
			fmt.Println("格納エラー",err)
			return
		} else {
            result = append(result, u)
        }
}
 fmt.Println(result)

 err = json.NewEncoder(w).Encode(result)
if err != nil {
	http.Error(w, "レスポンスのエンコードに失敗しました: "+err.Error(), http.StatusInternalServerError)
	return
}

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
	var persona model.Persona
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
  
		//ペルソナデータベースに挿入
	query := `  INSERT INTO persona (name, user_id, sex, age, profession, problems, behavior) VALUES ($1, $2, $3, $4, $5, $6, $7) returning id`
	var insertedID int
	err = db.QueryRow(query, persona.Name, 1, persona.Sex, persona.Age, persona.Profession, persona.Problems, persona.Behavior).Scan(&insertedID)
	if err != nil {
			http.Error(w, "データベースへの挿入に失敗しました: "+err.Error(), http.StatusInternalServerError)
			return
	}

	//コメントをAIから取得
	comment,err := util.CreatePersonaFirstComment(persona, openaiClient)
	
	//コメントをコメントDBに挿入
	query2 := `  INSERT INTO comment (comment) VALUES ($1) returning id`
	var insertedID2 int
	err2 := db.QueryRow(query2, comment).Scan(&insertedID2)
	if err2 != nil {
			http.Error(w, "データベースへの挿入に失敗しました: "+err.Error(), http.StatusInternalServerError)
			return
	}

	fmt.Println("データベースの挿入に成功しました")
	fmt.Fprintln(w, "登録が完了しました！")
}


