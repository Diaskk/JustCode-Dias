package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	dataForDb := "user=dias dbname=testdb sslmode=disable password=smth"
	db, err := sql.Open("postgres", dataForDb)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS table1 (id SERIAL PRIMARY KEY, name VARCHAR(50))")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("INSERT INTO people (name) VALUES ($1)", "Dias Kuan")
	if err != nil {
		log.Fatal(err)
	}
}
