package database

import (
"github.com/sashabaranov/go-openai"
"fmt"
"net/http"
"encoding/json"
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
