package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// ============================== DB variables ==============================

// MySQLDBHandler as type struct
type MySQLDBHandler struct {
	Conn *sqlx.DB
}

// User data struct for user table
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
	userTable      string = "users"
	postTable      string = "posts"
	commentTable   string = "comment"
)

// ============================== HTTP variables ==============================

// HTTPResponseVM base http viewmodel for http rest responses
type HTTPResponseVM struct {
	Status    int         `json:"-"`
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	ErrorCode interface{} `json:"errorCode,omitempty"`
	Data      interface{} `json:"data"`
}

// CreateUserRequest data struct for create user request
type CreateUserRequest struct {
	Email         string `json:"email"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	ContactNumber string `json:"contactNumber"`
}

// UpdateUserRequest data struct for update user request
type UpdateUserRequest struct {
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	ContactNumber string `json:"contactNumber"`
}

// UserResponse data struct for user response
type UserResponse struct {
	ID            int64  `json:"id"`
	Email         string `json:"email"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	ContactNumber string `json:"contactNumber"`
	CreatedAt     int64  `json:"createdAt"`
	UpdatedAt     int64  `json:"updatedAt"`
}

// PostRequest use for post request
type PostRequest struct {
	AuthorID int    `json:"authorId"`
	Content  string `json:"content"`
}

// PostResponse use for post response
type PostResponse struct {
	ID        int    `json:"id"`
	AuthorID  int    `json:"authorId"`
	Content   string `json:"content"`
	Timestamp int64  `json:"timestamp"`
}

// PatchRequest use for post request
type PatchRequest struct {
	Content string `json:"content"`
}

// PatchResponse use for post response
type PatchResponse struct {
	Content   string `json:"content"`
	Timestamp int64  `json:"timestamp"`
}

// StudentInformationResponse use for response student info
type StudentInformationResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	School string `json:"school"`
}

func main() {
	port := ":8080"
	fmt.Println("Starting Server....")

	// initialize mysql db handler
	mysqlDBHandler = &MySQLDBHandler{}

	// connect to database
	err := mysqlDBHandler.Connect("127.0.0.1", "3306", "nuxify_db_training", "root", "123")
	if err != nil {
		panic(err)
	}

	// initialize http router
	router := chi.NewRouter()

	// initialize middlewares
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Route("/api", func(router chi.Router) {
		router.Route("/post", func(router chi.Router) {
			router.Post("/", postHandler)
			router.Get("/{id}", getByIDHandler)
			router.Patch("/{id}", patchHandler)
			router.Delete("/{id}", deleteHandler)
		})

		router.Get("/posts", getAllHandler)

		router.Route("/user", func(router chi.Router) {
			router.Post("/", InsertUserHandler)
			router.Get("/{id}", GetUserByIDHandler)
			router.Patch("/{id}", UpdateUserHandler)
			router.Delete("/{id}", DeleteUserHandler)
		})

		router.Get("/users", GetAllUsersHandler)
	})

	fmt.Println("Server is listening on " + port)
	log.Fatal(http.ListenAndServe(port, router))
}

// ============================== HTTP methods ==============================

// InsertUserHandler creates a new user resource
func InsertUserHandler(w http.ResponseWriter, r *http.Request) {
	var request CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response := HTTPResponseVM{
			Status:  http.StatusUnprocessableEntity,
			Success: false,
			Message: "Invalid payload sent.",
		}

		response.JSON(w)
		return
	}

	// verify content must not empty
	if len(request.Email) == 0 || len(request.FirstName) == 0 || len(request.LastName) == 0 {
		response := HTTPResponseVM{
			Status:  http.StatusUnprocessableEntity,
			Success: false,
			Message: "User input cannot be empty.",
		}

		response.JSON(w)
		return
	}

	// insert to database
	user := &User{
		Email:         request.Email,
		FirstName:     request.FirstName,
		LastName:      request.LastName,
		ContactNumber: request.ContactNumber,
	}

	stmt := fmt.Sprintf("INSERT INTO %s (email,first_name,last_name,contact_number) VALUES (:email,:first_name,:last_name,:contact_number)", userTable)
	res, err := mysqlDBHandler.Execute(stmt, user)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			response := HTTPResponseVM{
				Status:  http.StatusConflict,
				Success: false,
				Message: "Duplicate email.",
			}

			response.JSON(w)
			return
		}

		response := HTTPResponseVM{
			Status:  http.StatusInternalServerError,
			Success: false,
			Message: err.Error(),
		}

		response.JSON(w)
		return
	}

	// get id
	id, err := res.LastInsertId()
	if err != nil {
		response := HTTPResponseVM{
			Status:  http.StatusInternalServerError,
			Success: false,
			Message: err.Error(),
		}

		response.JSON(w)
		return
	}

	response := &HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully created user.",
		Data: &UserResponse{
			ID:            id,
			Email:         user.Email,
			FirstName:     user.FirstName,
			LastName:      user.LastName,
			ContactNumber: user.ContactNumber,
			CreatedAt:     time.Now().Unix(),
			UpdatedAt:     time.Now().Unix(),
		},
	}

	response.JSON(w)
}

// GetUserByIDHandler get user by id
func GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idNum, err := strconv.Atoi(id)
	if err != nil {
		response := &HTTPResponseVM{
			Status:  http.StatusUnprocessableEntity,
			Success: false,
			Message: "Invalid payload sent.",
		}

		response.JSON(w)
		return
	}

	// get from database
	var users []User

	// prepare statement
	stmt := fmt.Sprintf("SELECT * FROM %s WHERE id=:id", userTable)
	err = mysqlDBHandler.Query(stmt, map[string]interface{}{
		"id": idNum,
	}, &users)
	if err != nil {
		response := HTTPResponseVM{
			Status:  http.StatusInternalServerError,
			Success: false,
			Message: err.Error(),
		}

		response.JSON(w)
		return
	} else if len(users) == 0 {
		response := HTTPResponseVM{
			Status:  http.StatusNotFound,
			Success: false,
			Message: "Cannot find user.",
		}

		response.JSON(w)
		return
	}

	response := &HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully get user.",
		Data: &UserResponse{
			ID:            users[0].ID,
			Email:         users[0].Email,
			FirstName:     users[0].FirstName,
			LastName:      users[0].LastName,
			ContactNumber: users[0].ContactNumber,
			CreatedAt:     users[0].CreatedAt.Unix(),
			UpdatedAt:     users[0].UpdatedAt.Unix(),
		},
	}

	response.JSON(w)
}

// GetAllUsersHandler get all users
func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []User

	// prepare statement
	stmt := fmt.Sprintf("SELECT * FROM %s", userTable)
	err := mysqlDBHandler.Query(stmt, map[string]interface{}{}, &users)
	if err != nil {
		response := HTTPResponseVM{
			Status:  http.StatusInternalServerError,
			Success: false,
			Message: err.Error(),
		}

		response.JSON(w)
		return
	} else if len(users) == 0 {
		response := HTTPResponseVM{
			Status:  http.StatusNotFound,
			Success: false,
			Message: "No users found.",
		}

		response.JSON(w)
		return
	}

	var usersResponse []UserResponse

	for _, user := range users {
		usersResponse = append(usersResponse, UserResponse{
			ID:            user.ID,
			Email:         user.Email,
			FirstName:     user.FirstName,
			LastName:      user.LastName,
			ContactNumber: user.ContactNumber,
			CreatedAt:     user.CreatedAt.Unix(),
			UpdatedAt:     user.UpdatedAt.Unix(),
		})
	}

	response := &HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully get all users.",
		Data:    usersResponse,
	}
	response.JSON(w)
}

// DeleteUserHandler delete user
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idNum, err := strconv.Atoi(id)
	if err != nil {
		response := &HTTPResponseVM{
			Status:  http.StatusUnprocessableEntity,
			Success: false,
			Message: "Invalid payload sent.",
		}

		response.JSON(w)
		return
	}

	// delete in database
	user := &User{
		ID: int64(idNum),
	}

	// prepare statement
	stmt := fmt.Sprintf("DELETE FROM %s WHERE id=:id", userTable)
	_, err = mysqlDBHandler.Execute(stmt, user)
	if err != nil {
		response := HTTPResponseVM{
			Status:  http.StatusInternalServerError,
			Success: false,
			Message: err.Error(),
		}

		response.JSON(w)
		return
	}

	response := &HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully deleted user data.",
	}

	response.JSON(w)
}

// UpdateUserHandler updates a user resource
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idNum, err := strconv.Atoi(id)
	if err != nil {
		response := &HTTPResponseVM{
			Status:  http.StatusUnprocessableEntity,
			Success: false,
			Message: "Invalid payload sent.",
		}

		response.JSON(w)
		return
	}

	var request UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response := HTTPResponseVM{
			Status:  http.StatusUnprocessableEntity,
			Success: false,
			Message: "Invalid payload sent.",
		}

		response.JSON(w)
		return
	}

	// verify content must not empty
	if len(request.FirstName) == 0 || len(request.LastName) == 0 {
		response := HTTPResponseVM{
			Status:  http.StatusUnprocessableEntity,
			Success: false,
			Message: "User input cannot be empty.",
		}

		response.JSON(w)
		return
	}

	// insert to database
	user := &User{
		ID:            int64(idNum),
		FirstName:     request.FirstName,
		LastName:      request.LastName,
		ContactNumber: request.ContactNumber,
	}

	stmt := fmt.Sprintf("UPDATE %s SET first_name=:first_name,last_name=:last_name,contact_number=:contact_number WHERE id=:id", userTable)
	_, err = mysqlDBHandler.Execute(stmt, user)
	if err != nil {
		response := HTTPResponseVM{
			Status:  http.StatusInternalServerError,
			Success: false,
			Message: err.Error(),
		}

		response.JSON(w)
		return
	}

	response := &HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully updated user.",
		Data: &UserResponse{
			ID:            user.ID,
			Email:         user.Email,
			FirstName:     user.FirstName,
			LastName:      user.LastName,
			ContactNumber: user.ContactNumber,
			UpdatedAt:     time.Now().Unix(),
		},
	}

	response.JSON(w)
}

// deleteHandler handle delete function
func deleteHandler(w http.ResponseWriter, r *http.Request) {
	response := &HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully deleted post data.",
	}

	response.JSON(w)
}

// getAllHandler handle getall function
func getAllHandler(w http.ResponseWriter, r *http.Request) {
	var students []StudentInformationResponse

	students = append(students, StudentInformationResponse{
		ID:     123,
		Name:   "Loed",
		School: "HCDC",
	})

	students = append(students, StudentInformationResponse{
		ID:     124,
		Name:   "Karl",
		School: "HCDC",
	})

	students = append(students, StudentInformationResponse{
		ID:     125,
		Name:   "GarzAlma",
		School: "HCDC",
	})

	response := &HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully fetched posts data.",
		Data:    students,
	}

	response.JSON(w)
}

// getByIDHandler handle get by id function
func getByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idNum, err := strconv.Atoi(id)
	if err != nil {
		response := &HTTPResponseVM{
			Status:  http.StatusUnprocessableEntity,
			Success: false,
			Message: "Invalid payload sent.",
		}

		response.JSON(w)
		return
	}

	response := &HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully fetched post data.",
		Data: &StudentInformationResponse{
			ID:     idNum,
			Name:   "Karl",
			School: "HCDC",
		},
	}

	response.JSON(w)
}

// patchHandler handle update function
func patchHandler(w http.ResponseWriter, r *http.Request) {
	var request PatchRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response := HTTPResponseVM{
			Status:  http.StatusUnprocessableEntity,
			Success: false,
			Message: "Invalid payload sent.",
		}

		response.JSON(w)
		return
	}

	// verify vontent must not empty
	if len(strings.TrimSpace(request.Content)) == 0 {
		response := HTTPResponseVM{
			Status:  http.StatusUnprocessableEntity,
			Success: false,
			Message: "Content cannot be empty.",
		}

		response.JSON(w)
		return
	}

	response := &HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully updated post.",
		Data: &PatchResponse{
			Content:   request.Content,
			Timestamp: time.Now().Unix(),
		},
	}

	response.JSON(w)
}

// postHandler handle create function
func postHandler(w http.ResponseWriter, r *http.Request) {
	var request PostRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response := HTTPResponseVM{
			Status:  http.StatusUnprocessableEntity,
			Success: false,
			Message: "Invalid payload sent.",
		}

		response.JSON(w)
		return
	}

	// verify content must not empty
	if len(strings.TrimSpace(request.Content)) == 0 {
		response := HTTPResponseVM{
			Status:  http.StatusUnprocessableEntity,
			Success: false,
			Message: "Content cannot be empty.",
		}

		response.JSON(w)
		return
	}

	response := &HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully created post.",
		Data: &PostResponse{
			ID:        1,
			AuthorID:  request.AuthorID,
			Content:   request.Content,
			Timestamp: time.Now().Unix(),
		},
	}

	response.JSON(w)
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

// ============================== HTTP Helper ==============================

// JSON converts http responsewriter to json
func (response *HTTPResponseVM) JSON(w http.ResponseWriter) {
	if response.Data == nil {
		response.Data = map[string]interface{}{}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Status)
	_ = json.NewEncoder(w).Encode(response)
}
