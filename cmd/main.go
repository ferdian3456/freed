package main

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
	"os"
)

func main() {
	db, err := sql.Open("pgx", os.Getenv("DB_URI"))
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Printf("Failed to ping to database:%v\n", err)
		os.Exit(1)
	}

	fmt.Println("Success to connect and ping database")
}
