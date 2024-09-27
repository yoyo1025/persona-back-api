module github.com/yoyo1025/persona-back-api

go 1.19

require (
	github.com/lib/pq v1.10.9
	github.com/sashabaranov/go-openai v1.30.3
)

replace github.com/yoyo1025/persona-back-api => ./ // ローカル開発用
