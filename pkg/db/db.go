package db

import (
	"database/sql"
	"log"
	"os"
)

var db *sql.DB

func Connect() {
	var err error
	db, err = sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatal(err)
	}
}

func Insert(value int) {
	var stm string = `INSERT INTO reading(val) VALUES ($1)`

	res, err := db.Exec(stm, value)
	if err != nil || res == nil {
		log.Fatal(err)
	}
}
