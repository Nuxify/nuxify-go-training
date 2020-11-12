package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// User of type struct
type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

var schema string = "CREATE TABLE `users` (	  	`id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,		  	`name` varchar(255) NOT NULL		)"

func main() {
	conn, err := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/story")
	if err != nil {
		panic(err)
	}
	conn.MustExec(schema)
	res, err := conn.Exec("INSERT INTO users (name) VALUES(\"Peter\")")
	if err != nil {
		panic(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created user with id:%d", id)
	var user User
	err = conn.Get(&user, "select * from users where id=?", id)
	if err != nil {
		panic(err)
	}
	_, err = conn.Exec("UPDATE users set name=\"John\" where id=?", id)
	if err != nil {
		panic(err)
	}
	_, err = conn.Exec("DELETE FROM users where id=?", id)
	if err != nil {
		panic(err)
	}

}
