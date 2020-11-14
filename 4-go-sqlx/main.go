package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MySQLDBHandler struct {
	Conn *sqlx.DB
}

// User as new type struct
type User struct {
	ID            int64
	Email         string
	FirstName     string    `db:"first_name"`
	LastName      string    `db:"last_name"`
	ContactNumber string    `db:"contact_number"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

var (
	mysqlDBHandler *MySQLDBHandler
	tableName      string = "users"
)

func main() {
	// initialize mysql db handler
	mysqlDBHandler = &MySQLDBHandler{}

	// connect to database
	err := mysqlDBHandler.Connect("127.0.0.1", "3306", "nuxify_training", "root", "1234")
	if err != nil {
		panic(err)
	}

	//InsertUser()
	//GetUser(3)
	GetUsers()

}

// ====================== User ========================

// InsertUser fucntion that create the user
func InsertUser() {
	user := &User{
		Email:         "dogie@alma.com",
		FirstName:     "Dogie",
		LastName:      "Alma",
		ContactNumber: "+639456042882",
	}

	// form the statement
	stmt := fmt.Sprintf("INSERT INTO %s (email,first_name,last_name,contact_number) VALUES (:email,:first_name,:last_name,:contact_number)", tableName)
	res, err := mysqlDBHandler.Execute(stmt, user)
	if err != nil {
		log.Println(err)

		if strings.Contains(err.Error(), "Duplicate entry") {
			log.Println("Duplicate user")
		} else {
			panic(err)
		}
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("[USER ID]", id)
}

// GetUser function that get user by id
func GetUser(userID int64) {
	var users []User

	// prepare statement
	stmt := fmt.Sprintf("SELECT * FROM %s WHERE id=:id", tableName)
	err := mysqlDBHandler.Query(stmt, map[string]interface{}{
		"id": userID,
	}, &users)
	if err != nil {
		panic(err)
	}

	fmt.Println(users[0].ID, users[0].Email)
}

// GetUsers function that get all users
func GetUsers() {
	var users []User

	// prepare statement
	stmt := fmt.Sprintf("SELECT * FROM %s", tableName)
	err := mysqlDBHandler.Query(stmt, map[string]interface{}{}, &users)
	if err != nil {
		panic(err)
	}

	for _, user := range users {
		fmt.Println(user.ID, user.Email)
	}
}

// ====================== Post ========================

// ====================== Comment ========================

// ============================== MySQL Helper ==============================

// Connect opens a new connection to the mysql interface
func (h *MySQLDBHandler) Connect(host, port, database, username, password string) error {
	conn, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, database))
	if err != nil {
		return err
	}

	h.Conn = conn

	err = conn.Ping()
	if err != nil {
		connErr := fmt.Errorf("[SERVER] Error connecting to the database! %s", err.Error())

		return connErr
	}

	fmt.Println("[SERVER] Database connected successfully")

	return nil
}

// Execute executes the mysql statement following NamedExec
// It requires a valid sql statement and its struct
func (h *MySQLDBHandler) Execute(stmt string, model interface{}) (sql.Result, error) {
	res, err := h.Conn.NamedExec(stmt, model)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Query selects rows given by the sql statement
// It requires the statement, the model to bind the statement, and the target bind model for the results
func (h *MySQLDBHandler) Query(qstmt string, model interface{}, bindModel interface{}) error {
	nstmt, err := h.Conn.PrepareNamed(qstmt)
	if err != nil {
		return err
	}
	defer nstmt.Close()

	err = nstmt.Select(bindModel, model)
	if err != nil {
		return err
	}

	return nil
}
