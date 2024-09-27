FROM golang:1.19.1-alpine

# 必要なパッケージをインストール
RUN apk update && apk add --no-cache git

# 作業ディレクトリを設定
WORKDIR /app

# ホスト側のgo.modとgo.sumをコピー
COPY go.mod go.sum ./

# 依存関係をダウンロード
RUN go mod download

# アプリケーションのソースコードをコピー
COPY . .

# アプリケーションをビルド
RUN go build -o main .

# ポートを公開
EXPOSE 4000

# コンテナ起動時に実行するコマンドを設定
CMD ["./main"]
