package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// User as new type struct
type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

func main() {
	db, err := sqlx.Connect("mysql", "root:1234@(localhost:3306)/test_sqlx")
	if err != nil {
		log.Fatalln(err)
	}

	user := []User{}

	db.Select(&user, "select * from users")

	// createUser()

	log.Println("users...")
	log.Println(user)

}

func createUser() {
	db, err := sqlx.Connect("mysql", "root:1234@(localhost:3306)/test_sqlx")
	if err != nil {
		panic(err)
	}
	res, err := db.Exec("INSERT INTO users (name) VALUES(\"Additional test user\")")
	if err != nil {
		panic(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created user with id:%d", id)
}
