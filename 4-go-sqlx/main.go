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

// MySQLDBHandler as type struct
type MySQLDBHandler struct {
	Conn *sqlx.DB
}

// User as new type struct
type User struct {
	ID            int64
	Email         string
	FirstName     string    `db:"first_name"`
	LastName      string    `db:"last_name"`
	ContactNumber string    `db:"contact"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

// Post as new type struct
type Post struct {
	ID        int64
	AuthorID  int64     `db:"author_id"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// Comment as new type struct
type Comment struct {
	ID        int64
	PostID    int64     `db:"post_id"`
	AuthorID  int64     `db:"author_id"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

var (
	mysqlDBHandler *MySQLDBHandler
	userTable      string = "users"
	postTable      string = "posts"
	commentTable   string = "comment"
)

func main() {
	// initialize mysql db handler
	mysqlDBHandler = &MySQLDBHandler{}

	// connect to database
	err := mysqlDBHandler.Connect("127.0.0.1", "3306", "nuxify_training", "root", "1234")
	if err != nil {
		panic(err)
	}

	// ========= User =========
	// InsertUser()
	// GetUser(3)
	// GetUsers()
	// DeleteUser(6)
	// UpdateUser(6)

	// ========= Post =========

	// InsertPost(6)
	// GetPost(6)
	// GetPosts()
	// DeletePost(6)
	// UpdatePost(3)

	// ========= Comment =========

	// InsertComment(6)
	// GetComment(6)
	// GetComments()
	// DeleteComment(3)
	// UpdateComment(2)
}

// ====================== User ========================

// InsertUser fucntion that create the user
func InsertUser() {
	user := &User{
		Email:         "5thexample@gmail.com",
		FirstName:     "SecondTest",
		LastName:      "User",
		ContactNumber: "+639456042882",
	}

	// form the statement
	stmt := fmt.Sprintf("INSERT INTO %s (email,first_name,last_name,contact) VALUES (:email,:first_name,:last_name,:contact)", userTable)
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
	stmt := fmt.Sprintf("SELECT * FROM %s WHERE id=:id", userTable)
	err := mysqlDBHandler.Query(stmt, map[string]interface{}{
		"id": userID,
	}, &users)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(users[0].ID, users[0].Email)
}

// GetUsers function that get all users
func GetUsers() {
	var users []User

	// prepare statement
	stmt := fmt.Sprintf("SELECT * FROM %s", userTable)
	err := mysqlDBHandler.Query(stmt, map[string]interface{}{}, &users)
	if err != nil {
		log.Println(err)
	}

	for _, user := range users {
		fmt.Println(user.ID, user.Email)
	}
}

// DeleteUser function that delete user
func DeleteUser(userID int64) {
	user := &User{
		ID: userID,
	}

	// form the statement
	stmt := fmt.Sprintf("DELETE %s WHERE id=:id", userTable)
	_, err := mysqlDBHandler.Execute(stmt, user)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("User successfully deleted..")
}

// UpdateUser function that update user
func UpdateUser(userID int64) {
	user := &User{
		ID:            userID,
		Email:         "updatedemail@example.com",
		FirstName:     "Lecty",
		LastName:      "Eisenach",
		ContactNumber: "+639456042882",
	}

	// form the statement
	stmt := fmt.Sprintf("UPDATE %s SET email=:email, first_name=:first_name, last_name=:last_name WHERE id=:id", userTable)
	_, err := mysqlDBHandler.Execute(stmt, user)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("User successfully updated..")
}

// ====================== Post ========================

// InsertPost create a new post function
func InsertPost(authorID int64) {
	post := &Post{
		AuthorID: authorID,
		Content:  "New post has been created",
	}

	// form the statement
	stmt := fmt.Sprintf("INSERT INTO %s (author_id, content) VALUES (:author_id, :content)", postTable)
	res, err := mysqlDBHandler.Execute(stmt, post)
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

	fmt.Println("[POST ID]", id)
}

// GetPost function that get user by id
func GetPost(postID int64) {
	var posts []Post

	// prepare statement
	stmt := fmt.Sprintf("SELECT * FROM %s WHERE id=:id", postTable)
	err := mysqlDBHandler.Query(stmt, map[string]interface{}{
		"author_id": postID,
	}, &posts)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(posts[0].ID, posts[0].Content)
}

//GetPosts Get all posts fucntion
func GetPosts() {
	var posts []Post

	// prepare statement
	stmt := fmt.Sprintf("SELECT * FROM %s", postTable)
	err := mysqlDBHandler.Query(stmt, map[string]interface{}{}, &posts)
	if err != nil {
		log.Println(err)
	}

	for _, post := range posts {
		fmt.Println(post.AuthorID, post.Content)
	}
}

// DeletePost function that delete user
func DeletePost(postID int64) {
	post := &Post{
		ID: postID,
	}

	// form the statement
	stmt := fmt.Sprintf("DELETE FROM %s WHERE id=:id", postTable)
	_, err := mysqlDBHandler.Execute(stmt, post)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Post successfully deleted..")
}

// UpdatePost function that update post
func UpdatePost(postID int64) {
	post := &Post{
		ID:      postID,
		Content: "Check post update",
	}

	// form the statement
	stmt := fmt.Sprintf("UPDATE %s SET content=:content WHERE id=:id", postTable)
	_, err := mysqlDBHandler.Execute(stmt, post)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Post successfully updated..")
}

// ====================== Comment ========================

// InsertComment create a new comment function
func InsertComment(authorID int64) {
	comment := &Comment{
		PostID:   3,
		AuthorID: authorID,
		Content:  "Blah Blah",
	}

	// form the statement
	stmt := fmt.Sprintf("INSERT INTO %s (post_id, author_id, content) VALUES (:post_id, :author_id, :content)", commentTable)
	res, err := mysqlDBHandler.Execute(stmt, comment)
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

	fmt.Println("[COMMENT ID]", id)
}

// GetComment function that get comment by id
func GetComment(commentID int64) {
	var comments []Comment

	// prepare statement
	stmt := fmt.Sprintf("SELECT * FROM %s WHERE id=:id", commentTable)
	err := mysqlDBHandler.Query(stmt, map[string]interface{}{
		"author_id": commentID,
	}, &comments)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(comments[0].ID, comments[0].Content)
}

//GetComments Get all comments fucntion
func GetComments() {
	var comments []Comment

	// prepare statement
	stmt := fmt.Sprintf("SELECT * FROM %s", commentTable)
	err := mysqlDBHandler.Query(stmt, map[string]interface{}{}, &comments)
	if err != nil {
		log.Println(err)
	}

	for _, post := range comments {
		fmt.Println(post.AuthorID, post.Content)
	}
}

// DeleteComment function that delete comment
func DeleteComment(commentID int64) {
	comment := &Comment{
		ID: commentID,
	}

	// form the statement
	stmt := fmt.Sprintf("DELETE FROM %s WHERE id=:id", commentTable)
	_, err := mysqlDBHandler.Execute(stmt, comment)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Comment successfully deleted..")
}

// UpdateComment function that update comment
func UpdateComment(commentID int64) {
	comment := &Comment{
		ID:      commentID,
		Content: "Update Comment",
	}

	// form the statement
	stmt := fmt.Sprintf("UPDATE %s SET content=:content WHERE id=:id", commentTable)
	_, err := mysqlDBHandler.Execute(stmt, comment)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Comment successfully updated..")
}

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
