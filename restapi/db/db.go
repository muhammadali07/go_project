package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func GetDB() *sqlx.DB {
	return db
}

func init() {
	var err error
	db, err = sqlx.Open("postgres", "host=godb port=5432 user=gorestuser password=gorestpass dbname=gorestdevelopment sslmode=disable")
	if err != nil {
		panic(err)
	}
}
