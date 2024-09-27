package database

import "database/sql"

// GetDB は他のパッケージからデータベース接続を取得するための関数
func GetDB() *sql.DB {
	return db
}
