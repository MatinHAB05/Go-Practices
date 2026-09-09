package main

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	conn, err := sql.Open("mysql", "testapp:testpass@tcp(:3308)/testdb")
	if err != nil {
		log.Fatal(err)
	}

	queries := db.New(conn)

	user, err := queries.CreateUser(context.Background(), db.CreateUserParams{
		Name:  "Ali",
		Email: "ali@test.com",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(user)
}
