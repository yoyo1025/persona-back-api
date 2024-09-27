FROM golang:1.19.1-alpine

# 必要なパッケージをインストール
RUN apk update && apk add --no-cache git bash

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

# wait-for-it.shをコピーして、Postgresの起動を待機できるようにする
COPY wait-for-it.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/wait-for-it.sh

# ポートを公開
EXPOSE 4000

CMD ["wait-for-it.sh", "persona_db:5432", "--", "./main"]
