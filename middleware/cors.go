// corsMiddleware.go
package middleware

import (
	"net/http"
)

func CORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // 必要なCORSヘッダーを設定
        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5200")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        // OPTIONSメソッドの場合、ここでリターン
        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusOK)
            return
        }

        // 次のハンドラーを実行
        next.ServeHTTP(w, r)
    })
}
