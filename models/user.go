package models

import (
	"database/sql"
	"log"
)

type User struct {
	id   int
	name string
}

func AllUsers(db *sql.DB) *sql.Rows {
	rows, err := db.Query("SELECT * FROM `users`")
	if err != nil {
		log.Fatal("エラー：", err)
		// なんか返す
	}
	return rows
}
