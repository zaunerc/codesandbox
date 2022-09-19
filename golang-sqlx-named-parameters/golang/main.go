// file golang-sqlx-named-parameters/golang/main.go
package main

import (
	"fmt"
	"log"
	"main/model"
	"main/repository/bookrepo"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

func main() {
	db := MustCreateDBConn()

	book := &model.Book{
		Title: "alice in wonderland",
	}

	createdBook, err := bookrepo.Create(db, book)
	if err != nil {
		log.Fatalf("Error while creating new book: %s", err.Error())
	}
	log.Printf("Created book: %+v", createdBook)

	log.Printf("Now going to search for books...")

	bookSearchParams := &bookrepo.SearchParams{
		Title: func() *string { val := "alice in wonderland"; return &val }(),
	}
	foundBooks, err := bookrepo.Search(db, bookSearchParams)
	if err != nil {
		log.Fatalf("Error while searching for new book: %s", err.Error())
	}
	log.Printf("Found books: %+v", foundBooks)

	log.Printf("Finished successfully")
}

func MustCreateDBConn() *sqlx.DB {

	host := "localhost"
	port := 5432
	username := "postgres"
	password := "postgrespwd"
	dbName := "postgres"

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbName)
	log.Printf("Creating DB conn host:%s port:%v username:%s", host, port, username)
	db := sqlx.MustOpen("postgres", psqlInfo)

	// db.MapperFunc can be used to customize how SQL
	// named params are mapped to Golang fields.

	/*
		db.MapperFunc(func(name string) string {
			return name
		})
	*/

	log.Printf("DB conn created")
	return db
}
