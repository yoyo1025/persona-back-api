package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB // グローバル変数としてデータベース接続を保持

func InitDB() {
	var err error
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
	)
	db, err = sql.Open("postgres", connStr) // ローカル変数ではなく、グローバル変数に代入
	if err != nil {
		log.Fatal("データベースへの接続に失敗しました:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("データベースへの接続確認に失敗しました:", err)
	}
}