package router

import (
	"fmt"
	"net/http"
	"github.com/yoyo1025/persona-back-api/database"
	"github.com/sashabaranov/go-openai"
	"github.com/yoyo1025/persona-back-api/middleware"
)

func NewRouter(openaiClient *openai.Client) http.Handler {
    mux := http.NewServeMux()
		mux.HandleFunc("/test", testHandler)
		mux.HandleFunc("/", database.UseridHandler)
    // ミドルウェアの適用
    return middleware.CORS(mux)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, "GET called!")
	} else if r.Method == "POST" {
		fmt.Fprintf(w, "POST called")
	}
}