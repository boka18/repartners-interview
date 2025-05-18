package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/boka18/repartners-interview/handlers"
)

func connectToDB() *sql.DB {
	env := os.Getenv("ENV")
	sslMode := "disable"
	if env == "develop" {
		sslMode = "require"
	}
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		sslMode,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	return db
}

func main() {
	db := connectToDB()
	err := db.Ping()
	if err != nil {
		fmt.Printf("err pinging db: %s", err.Error())
	}
	defer db.Close()

	mux := http.NewServeMux()
	handlers.RegisterStaticHandler(mux)
	handlers.RegisterCalculateHandler(mux, db)
	handlers.RegisterPackSizeHandler(mux, db)

	fmt.Println("Listening on http://localhost:80")
	log.Fatal(http.ListenAndServe(":80", mux))
}
