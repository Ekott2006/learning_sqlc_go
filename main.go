package main

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"learning_sqlc/tutorial"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
var schema string

func main() {
	fmt.Println(schema)

	ctx := context.Background()

	db, err := sql.Open("sqlite3", "./author.sqlite3")
	if err != nil {
		log.Fatal(err)
		return
	}

	if _, err := db.ExecContext(ctx, schema); err != nil {
		log.Fatal(err)
		return
	}

	queries := tutorial.New(db)

	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(authors)

	insertedAuthor, err := queries.CreateAuthor(ctx, tutorial.CreateAuthorParams{Name: "Nsikak Ekott", Bio: sql.NullString{String: "A Fresh Programmer looking for money"}})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(insertedAuthor)

	updatedAuthor, err := queries.UpdateAuthor(ctx, tutorial.UpdateAuthorParams{ID: insertedAuthor.ID, Bio: insertedAuthor.Bio, Name: "Ubong Ekott"})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(updatedAuthor)
	getAuthor, err := queries.GetAuthor(ctx, insertedAuthor.ID)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(getAuthor)
}
