package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func GetDB() *sqlx.DB {
	return db
}

func init() {

	env := godotenv.Load()
	if env != nil {
		fmt.Println("Error loading .env file")
		return
	}

	// Baca variabel lingkungan dari file .env
	dbHost := os.Getenv("postgres_db_host")
	dbPort := os.Getenv("postgres_db_port")
	dbUser := os.Getenv("postgres_db_user")
	dbPassword := os.Getenv("postgres_db_password")
	dbName := os.Getenv("postgres_db")

	var err error
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err = sqlx.Open("postgres", conn)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Database connection established so far")
	}

	// Membuat tabel users jika belum ada
	db.MustExec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT,
		created_at TIMESTAMP

	)`)
}
