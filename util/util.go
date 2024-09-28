package util

import (
 "context"
 "encoding/json"
 "fmt"
 "net/http"

 "github.com/sashabaranov/go-openai"
 "github.com/yoyo1025/persona-back-api/model" 
)

func CreatePersonaFirstComment(persona model.Persona, client *openai.Client) (string, error) {
 prompt := fmt.Sprintf(
  "次の情報は架空の困りごとをペルソナです。あなたはそのペルソナのふりをしてください。ペルソナの設定はその都度付け加えても構いません。次の情報から軽く自己紹介をしてください。改行は要りません。:\n"+
   "名前: %s\n性別: %s\n年齢: %d\n職業: %s\n問題: %s\n行動: %s\n",
  persona.Name, persona.Sex, persona.Age, persona.Profession, persona.Problems, persona.Behavior,
 )

 resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
  Model: openai.GPT4,
  Messages: []openai.ChatCompletionMessage{
   {
    Role:    openai.ChatMessageRoleUser,
    Content: prompt,
   },
  },
 })

 if err != nil {
  return "", fmt.Errorf("OpenAI APIエラー: %v", err)
 }

 return resp.Choices[0].Message.Content, nil
}

func CreateComment(comments []model.Comment, client *openai.Client) (string, error) {
 // 会話履歴を展開
 var conversationHistory string
 for _, comment := range comments {
  conversationHistory += fmt.Sprintf("%s\n", comment.Comment)
 }

 // プロンプトを作成
 prompt := fmt.Sprintf(
  "今から渡す文章のペルソナになり切ったつもりで、次の会話履歴の流れに沿うように簡潔に返答してください。改行は要りません。\n会話履歴:\n%s",
  conversationHistory,
 )

 // OpenAIのChatCompletion APIを呼び出して応答を生成
 resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
  Model: openai.GPT4,
  Messages: []openai.ChatCompletionMessage{
   {
    Role:    openai.ChatMessageRoleUser,
    Content: prompt,
   },
  },
 })

 if err != nil {
  return "", fmt.Errorf("OpenAI APIエラー: %v", err)
 }

 return resp.Choices[0].Message.Content, nil
}

// AIに要件定義書を作成させる関数
func CreateDocument(w http.ResponseWriter, r *http.Request, client *openai.Client) {
 fmt.Println("要件定義書を作成する用の関数が呼び出されました")

 // POSTメソッド以外のリクエストを拒否
 if r.Method != http.MethodPost {
  http.Error(w, "許可されていないメソッドです", http.StatusMethodNotAllowed)
  return
 }

 // リクエストボディをデコード
 var comments []model.Comment
 err := json.NewDecoder(r.Body).Decode(&comments)
 if err != nil {
  http.Error(w, "リクエストデータのデコードに失敗しました", http.StatusBadRequest)
  return
 }

 fmt.Println("ドキュメント作成開始...")
 // コメント履歴をAIに渡して要件定義書を生成
 document, err := GenerateRequirementsDocument(comments, client)
 if err != nil {
  http.Error(w, fmt.Sprintf("AIによる要件定義書作成に失敗しました: %v", err), http.StatusInternalServerError)
  return
 }
 fmt.Println(document)
 

 // AIが作成した要件定義書をJSONで返す
 w.Header().Set("Content-Type", "application/json")
 w.WriteHeader(http.StatusOK)
 err = json.NewEncoder(w).Encode(map[string]string{"document": document})
 if err != nil {
  http.Error(w, "レスポンスのエンコードに失敗しました: "+err.Error(), http.StatusInternalServerError)
 }
}

// AIに要件定義書を作成させるための関数
func GenerateRequirementsDocument(comments []model.Comment, client *openai.Client) (string, error) {
 // 会話履歴を展開
 var conversationHistory string
 for _, comment := range comments {
  conversationHistory += fmt.Sprintf("%s\n", comment.Comment)
 }

 // プロンプトを作成
 prompt := fmt.Sprintf(
  "次の会話履歴をもとに、要件定義書(主に機能要件と非機能要件)を作成してください。要件定義書には主な機能やユーザーストーリー、必要なAPIや技術的な要件を含めてください。マークダウン形式にしてください。:\n\n%s",
  conversationHistory,
 )

 // OpenAIのChatCompletion APIを呼び出して応答を生成
 resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
  Model: openai.GPT4o,
  Messages: []openai.ChatCompletionMessage{
   {
    Role:    openai.ChatMessageRoleUser,
    Content: prompt,
   },
  },
 })

 if err != nil {
  return "", fmt.Errorf("OpenAI APIエラー: %v", err)
 }

 // 応答を返す
 return resp.Choices[0].Message.Content, nil
}
