package main

import (
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

	log.Println("users...")
	log.Println(user)

}
