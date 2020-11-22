package main

import (
	"database/sql"
	"encoding/json"
	"errors"
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

// Post data struct for post table
type Post struct {
	ID        int64
	AuthorID  int64 `db:"author_id"`
	Content   string
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// Comment data struct for post table
type Comment struct {
	ID        int64
	PostID    int64 `db:"post_id"`
	AuthorID  int64 `db:"author_id"`
	Content   string
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
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

// ============================== User HTTP variables ==============================

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

// ============================== Post HTTP variables ==============================

// PostResponse use for post response
type PostResponse struct {
	ID        int64  `json:"id"`
	Author    Author `json:"author"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

// Author response struct for author
type Author struct {
	ID            int64  `json:"id"`
	Email         string `json:"email"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	ContactNumber string `json:"contactNumber"`
	CreatedAt     int64  `json:"createdAt"`
	UpdatedAt     int64  `json:"updatedAt"`
}

// CreatePostRequest use for post request
type CreatePostRequest struct {
	AuthorID int64  `json:"authorId"`
	Content  string `json:"content"`
}

// UpdatePostRequest use for post request
type UpdatePostRequest struct {
	AuthorID int64  `json:"authorId"`
	Content  string `json:"content"`
}

// ============================== Comment HTTP variables ==============================

// CommentResponse use for post response
type CommentResponse struct {
	ID        int64  `json:"id"`
	PostID    int64  `json:"postId"`
	Author    Author `json:"author"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

// CreateCommentRequest use for post request
type CreateCommentRequest struct {
	PostID   int64  `json:"postId"`
	AuthorID int64  `json:"authorId"`
	Content  string `json:"content"`
}

// UpdateCommentRequest use for post request
type UpdateCommentRequest struct {
	AuthorID int64  `json:"authorId"`
	Content  string `json:"content"`
}

// initialize main function
func main() {
	port := ":8080"
	fmt.Println("Starting Server....")

	// initialize mysql db handler
	mysqlDBHandler = &MySQLDBHandler{}

	// connect to database
	err := mysqlDBHandler.Connect("127.0.0.1", "3306", "nuxify_training", "root", "1234")
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
		// routes for user
		router.Route("/user", func(router chi.Router) {
			router.Post("/", CreateUserHandler)
			router.Get("/{id}", GetUserByIDHandler)
			router.Patch("/{id}", UpdateUserHandler)
			router.Delete("/{id}", DeleteUserHandler)
		})

		router.Get("/users", GetAllUsersHandler)

		// routes for post
		router.Route("/post", func(router chi.Router) {
			router.Post("/", CreatePostHandler)
			router.Get("/{id}", GetPostByIDHandler)
			router.Get("/{postID}/comments", GetCommentsByPostIDHandler)
			router.Patch("/{id}", UpdatePostHandler)
			router.Delete("/{id}", DeletePostHandler)
		})

		router.Get("/posts", GetAllPostsHandler)

		// routes for comment
		router.Route("/comment", func(router chi.Router) {
			router.Post("/", CreateCommentHandler)
			router.Get("/{id}", GetCommentByIDHandler)
			router.Patch("/{id}", UpdateCommentHandler)
			router.Delete("/{id}", DeleteCommenttHandler)
		})

		router.Get("/comments", GetAllCommentsHandler)
	})

	fmt.Println("Server is listening on " + port)
	log.Fatal(http.ListenAndServe(port, router))
}

// ============================== HTTP methods ==============================
// ============================== users handler ==============================

// CreateUserHandler creates a new user resource
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
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

	user := User{
		Email:         request.Email,
		FirstName:     request.FirstName,
		LastName:      request.LastName,
		ContactNumber: request.ContactNumber,
	}

	// insert to database
	id, err := InsertUserRepository(user)
	if err != nil {
		if err.Error() == "DUPLICATE_EMAIL" {
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

	// get user
	user, err := SelectUserByIDRepository(int64(idNum))
	if err != nil {
		if err.Error() == "MISSING_RECORD" {
			response := HTTPResponseVM{
				Status:  http.StatusNotFound,
				Success: false,
				Message: "Cannot find user.",
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

	response := &HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully get user.",
		Data: &UserResponse{
			ID:            user.ID,
			Email:         user.Email,
			FirstName:     user.FirstName,
			LastName:      user.LastName,
			ContactNumber: user.ContactNumber,
			CreatedAt:     user.CreatedAt.Unix(),
			UpdatedAt:     user.UpdatedAt.Unix(),
		},
	}

	response.JSON(w)
}

// GetAllUsersHandler get all users
func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {

	// get from database
	users, err := SelectUsersRepository()
	if err != nil {
		if err.Error() == "MISSING_RECORD" {
			response := HTTPResponseVM{
				Status:  http.StatusNotFound,
				Success: false,
				Message: "Cannot find user.",
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
	user := User{
		ID:            int64(idNum),
		FirstName:     request.FirstName,
		LastName:      request.LastName,
		ContactNumber: request.ContactNumber,
	}

	_, err = UpdateUserRepository(user)
	if err != nil {
		if err.Error() == "MISSING_RECORD" {
			response := HTTPResponseVM{
				Status:  http.StatusNotFound,
				Success: false,
				Message: "Cannot find user.",
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
	user := User{
		ID: int64(idNum),
	}

	// prepare statement
	err = DeleteUserRepository(user)
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

// ============================== posts handler ==============================

// CreatePostHandler handle create function
func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	var request CreatePostRequest
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
	if len(request.Content) == 0 {
		response := HTTPResponseVM{
			Status:  http.StatusUnprocessableEntity,
			Success: false,
			Message: "Post input cannot be empty.",
		}

		response.JSON(w)
		return
	}

	// insert to database
	post := Post{
		AuthorID: request.AuthorID,
		Content:  request.Content,
	}

	id, err := InsertPostRepository(post)
	if err != nil {
		response := HTTPResponseVM{
			Status:  http.StatusInternalServerError,
			Success: false,
			Message: err.Error(),
		}

		response.JSON(w)
		return
	}

	// get author details
	user, err := SelectUserByIDRepository(post.AuthorID)
	if err != nil {
		if err.Error() == "MISSING_RECORD" {
			response := HTTPResponseVM{
				Status:  http.StatusNotFound,
				Success: false,
				Message: "Cannot find user.",
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

	response := &HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully get post.",
		Data: &PostResponse{
			ID: id,
			Author: Author{
				ID:            user.ID,
				Email:         user.Email,
				FirstName:     user.FirstName,
				LastName:      user.LastName,
				ContactNumber: user.ContactNumber,
				CreatedAt:     user.CreatedAt.Unix(),
				UpdatedAt:     user.UpdatedAt.Unix(),
			},
			Content:   post.Content,
			CreatedAt: post.CreatedAt.Unix(),
			UpdatedAt: post.UpdatedAt.Unix(),
		},
	}

	response.JSON(w)
}

// GetPostByIDHandler get post by id
func GetPostByIDHandler(w http.ResponseWriter, r *http.Request) {
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

	// prepare statement
	post, err := SelectPostByIDRepository(int64(idNum))
	if err != nil {
		if err.Error() == "MISSING_RECORD" {
			response := HTTPResponseVM{
				Status:  http.StatusNotFound,
				Success: false,
				Message: "Cannot find user.",
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

	// get author details
	user, err := SelectUserByIDRepository(post.AuthorID)
	if err != nil {
		if err.Error() == "MISSING_RECORD" {
			response := HTTPResponseVM{
				Status:  http.StatusNotFound,
				Success: false,
				Message: "Cannot find user.",
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

	response := &HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully get post.",
		Data: &PostResponse{
			ID: post.ID,
			Author: Author{
				ID:            user.ID,
				Email:         user.Email,
				FirstName:     user.FirstName,
				LastName:      user.LastName,
				ContactNumber: user.ContactNumber,
				CreatedAt:     user.CreatedAt.Unix(),
				UpdatedAt:     user.UpdatedAt.Unix(),
			},
			Content:   post.Content,
			CreatedAt: post.CreatedAt.Unix(),
			UpdatedAt: post.UpdatedAt.Unix(),
		},
	}

	response.JSON(w)
}

// GetAllPostsHandler get all users
func GetAllPostsHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := SelectPostsRepository()
	if err != nil {
		if err.Error() == "MISSING_RECORD" {
			response := HTTPResponseVM{
				Status:  http.StatusNotFound,
				Success: false,
				Message: "Cannot find posts.",
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

	var postsResponse []PostResponse

	for _, post := range posts {
		// get author details
		user, err := SelectUserByIDRepository(post.AuthorID)
		if err != nil {
			if err.Error() == "MISSING_RECORD" {
				response := HTTPResponseVM{
					Status:  http.StatusNotFound,
					Success: false,
					Message: "Cannot find user.",
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
		postsResponse = append(postsResponse, PostResponse{
			ID: post.ID,
			Author: Author{
				ID:            user.ID,
				Email:         user.Email,
				FirstName:     user.FirstName,
				LastName:      user.LastName,
				ContactNumber: user.ContactNumber,
				CreatedAt:     user.CreatedAt.Unix(),
				UpdatedAt:     user.UpdatedAt.Unix(),
			},
			Content:   post.Content,
			CreatedAt: post.CreatedAt.Unix(),
			UpdatedAt: post.UpdatedAt.Unix(),
		})
	}

	response := &HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully get all posts.",
		Data:    postsResponse,
	}
	response.JSON(w)
}

// UpdatePostHandler updates a user resource
func UpdatePostHandler(w http.ResponseWriter, r *http.Request) {
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

	var request UpdatePostRequest
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
	if len(request.Content) == 0 {
		response := HTTPResponseVM{
			Status:  http.StatusUnprocessableEntity,
			Success: false,
			Message: "User input cannot be empty.",
		}

		response.JSON(w)
		return
	}

	// update to database
	post := Post{
		ID:       int64(idNum),
		AuthorID: request.AuthorID,
		Content:  request.Content,
	}

	post, err = UpdatePostRepository(post)
	if err != nil {
		if err.Error() == "MISSING_RECORD" {
			response := HTTPResponseVM{
				Status:  http.StatusNotFound,
				Success: false,
				Message: "Cannot find post.",
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

	// get author details
	user, err := SelectUserByIDRepository(post.AuthorID)
	if err != nil {
		if err.Error() == "MISSING_RECORD" {
			response := HTTPResponseVM{
				Status:  http.StatusNotFound,
				Success: false,
				Message: "Cannot find user.",
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

	response := &HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully get post.",
		Data: &PostResponse{
			ID: post.ID,
			Author: Author{
				ID:            user.ID,
				Email:         user.Email,
				FirstName:     user.FirstName,
				LastName:      user.LastName,
				ContactNumber: user.ContactNumber,
				CreatedAt:     user.CreatedAt.Unix(),
				UpdatedAt:     user.UpdatedAt.Unix(),
			},
			Content:   post.Content,
			UpdatedAt: post.UpdatedAt.Unix(),
		},
	}

	response.JSON(w)
}

// DeletePostHandler delete user
func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
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
	post := Post{
		ID: int64(idNum),
	}

	err = DeletePostRepository(post)
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
		Message: "Successfully deleted post data.",
	}

	response.JSON(w)
}

// ============================== comment handler ==============================

// CreateCommentHandler handle create function
func CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
	var request CreateCommentRequest
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
	if len(request.Content) == 0 {
		response := HTTPResponseVM{
			Status:  http.StatusUnprocessableEntity,
			Success: false,
			Message: "Post input cannot be empty.",
		}

		response.JSON(w)
		return
	}

	// insert to database
	comment := Comment{
		PostID:   request.PostID,
		AuthorID: request.AuthorID,
		Content:  request.Content,
	}

	id, err := InsertCommentRepository(comment)
	if err != nil {
		response := HTTPResponseVM{
			Status:  http.StatusInternalServerError,
			Success: false,
			Message: err.Error(),
		}

		response.JSON(w)
		return
	}

	// get author details
	user, err := SelectUserByIDRepository(comment.AuthorID)
	if err != nil {
		if err.Error() == "MISSING_RECORD" {
			response := HTTPResponseVM{
				Status:  http.StatusNotFound,
				Success: false,
				Message: "Cannot find user.",
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

	response := &HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully get post.",
		Data: &PostResponse{
			ID: id,
			Author: Author{
				ID:            user.ID,
				Email:         user.Email,
				FirstName:     user.FirstName,
				LastName:      user.LastName,
				ContactNumber: user.ContactNumber,
				CreatedAt:     user.CreatedAt.Unix(),
				UpdatedAt:     user.UpdatedAt.Unix(),
			},
			Content:   comment.Content,
			CreatedAt: comment.CreatedAt.Unix(),
			UpdatedAt: comment.UpdatedAt.Unix(),
		},
	}

	response.JSON(w)
}

// GetCommentByIDHandler get post by id
func GetCommentByIDHandler(w http.ResponseWriter, r *http.Request) {
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
	comment, err := SelectCommentByIDRepository(int64(idNum))
	if err != nil {
		if err.Error() == "MISSING_RECORD" {
			response := HTTPResponseVM{
				Status:  http.StatusNotFound,
				Success: false,
				Message: "Cannot find user.",
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

	// get author details
	user, err := SelectUserByIDRepository(comment.AuthorID)
	if err != nil {
		if err.Error() == "MISSING_RECORD" {
			response := HTTPResponseVM{
				Status:  http.StatusNotFound,
				Success: false,
				Message: "Cannot find user.",
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

	response := &HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully get comment.",
		Data: &CommentResponse{
			ID:     comment.ID,
			PostID: comment.PostID,
			Author: Author{
				ID:            user.ID,
				Email:         user.Email,
				FirstName:     user.FirstName,
				LastName:      user.LastName,
				ContactNumber: user.ContactNumber,
				CreatedAt:     user.CreatedAt.Unix(),
				UpdatedAt:     user.UpdatedAt.Unix(),
			},
			Content:   comment.Content,
			CreatedAt: comment.CreatedAt.Unix(),
			UpdatedAt: comment.UpdatedAt.Unix(),
		},
	}

	response.JSON(w)
}

// GetAllCommentsHandler get all users
func GetAllCommentsHandler(w http.ResponseWriter, r *http.Request) {
	// get from database
	comments, err := SelectCommentsRepository()
	if err != nil {
		if err.Error() == "MISSING_RECORD" {
			response := HTTPResponseVM{
				Status:  http.StatusNotFound,
				Success: false,
				Message: "Cannot find posts.",
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

	var commentResponse []CommentResponse

	for _, comment := range comments {
		// get author details
		user, err := SelectUserByIDRepository(comment.AuthorID)
		if err != nil {
			if err.Error() == "MISSING_RECORD" {
				response := HTTPResponseVM{
					Status:  http.StatusNotFound,
					Success: false,
					Message: "Cannot find user.",
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
		commentResponse = append(commentResponse, CommentResponse{
			ID:     comment.ID,
			PostID: comment.PostID,
			Author: Author{
				ID:            user.ID,
				Email:         user.Email,
				FirstName:     user.FirstName,
				LastName:      user.LastName,
				ContactNumber: user.ContactNumber,
				CreatedAt:     user.CreatedAt.Unix(),
				UpdatedAt:     user.UpdatedAt.Unix(),
			},
			Content:   comment.Content,
			CreatedAt: comment.CreatedAt.Unix(),
			UpdatedAt: comment.UpdatedAt.Unix(),
		})
	}

	response := &HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully get all comments.",
		Data:    commentResponse,
	}
	response.JSON(w)
}

// GetCommentsByPostIDHandler get comments by post id
func GetCommentsByPostIDHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "postID")
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
	comments, err := SelectCommentsByPostIDRepository(int64(idNum))
	if err != nil {
		if err.Error() == "MISSING_RECORD" {
			response := HTTPResponseVM{
				Status:  http.StatusNotFound,
				Success: false,
				Message: "Cannot find posts.",
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

	var commentResponse []CommentResponse

	for _, comment := range comments {
		// get user
		user, err := SelectUserByIDRepository(comment.AuthorID)
		if err != nil {
			if err.Error() == "MISSING_RECORD" {
				response := HTTPResponseVM{
					Status:  http.StatusNotFound,
					Success: false,
					Message: "Cannot find user.",
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

		commentResponse = append(commentResponse, CommentResponse{
			ID:      comment.ID,
			PostID:  comment.PostID,
			Content: comment.Content,
			Author: Author{
				ID:            user.ID,
				Email:         user.Email,
				FirstName:     user.FirstName,
				LastName:      user.LastName,
				ContactNumber: user.ContactNumber,
				CreatedAt:     user.CreatedAt.Unix(),
				UpdatedAt:     user.UpdatedAt.Unix(),
			},
			CreatedAt: comment.CreatedAt.Unix(),
			UpdatedAt: comment.UpdatedAt.Unix(),
		})
	}

	response := &HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully get all comments from post.",
		Data:    commentResponse,
	}
	response.JSON(w)
}

// UpdateCommentHandler updates a user resource
func UpdateCommentHandler(w http.ResponseWriter, r *http.Request) {
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

	var request UpdateCommentRequest
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
	if len(request.Content) == 0 {
		response := HTTPResponseVM{
			Status:  http.StatusUnprocessableEntity,
			Success: false,
			Message: "Comment input cannot be empty.",
		}

		response.JSON(w)
		return
	}

	// insert to database
	comment := Comment{
		ID:       int64(idNum),
		AuthorID: request.AuthorID,
		Content:  request.Content,
	}

	comment, err = UpdateCommentRepository(comment)
	if err != nil {
		if err.Error() == "MISSING_RECORD" {
			response := HTTPResponseVM{
				Status:  http.StatusNotFound,
				Success: false,
				Message: "Cannot find posts.",
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

	// get author details
	user, err := SelectUserByIDRepository(comment.AuthorID)
	if err != nil {
		if err.Error() == "MISSING_RECORD" {
			response := HTTPResponseVM{
				Status:  http.StatusNotFound,
				Success: false,
				Message: "Cannot find user.",
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

	response := &HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully updated comment.",
		Data: &CommentResponse{
			ID:     comment.ID,
			PostID: comment.PostID,
			Author: Author{
				ID:            user.ID,
				Email:         user.Email,
				FirstName:     user.FirstName,
				LastName:      user.LastName,
				ContactNumber: user.ContactNumber,
				CreatedAt:     user.CreatedAt.Unix(),
				UpdatedAt:     user.UpdatedAt.Unix(),
			},
			Content:   comment.Content,
			UpdatedAt: time.Now().Unix(),
		},
	}

	response.JSON(w)
}

// DeleteCommenttHandler delete user
func DeleteCommenttHandler(w http.ResponseWriter, r *http.Request) {
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
	comment := Comment{
		ID: int64(idNum),
	}

	// prepare statement
	err = DeleteCommentRepository(comment)
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
		Message: "Successfully deleted comment data.",
	}

	response.JSON(w)
}

// ============================== Repositories ==============================
// ============================== users repository ==============================

// InsertUserRepository insert a user data
func InsertUserRepository(data User) (int64, error) {
	stmt := fmt.Sprintf("INSERT INTO %s (email,first_name,last_name,contact_number) VALUES (:email,:first_name,:last_name,:contact_number)", userTable)
	res, err := mysqlDBHandler.Execute(stmt, data)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return -1, errors.New("DUPLICATE_EMAIL")
		}

		return -1, errors.New("DATABASE_ERROR")
	}

	// get id
	id, err := res.LastInsertId()
	if err != nil {
		return -1, errors.New("DATABASE_ERROR")
	}

	return id, nil
}

// SelectUserByIDRepository select user data by id
func SelectUserByIDRepository(ID int64) (User, error) {
	var users []User

	stmt := fmt.Sprintf("SELECT * FROM %s WHERE id=:id", userTable)
	err := mysqlDBHandler.Query(stmt, map[string]interface{}{
		"id": ID,
	}, &users)
	if err != nil {
		return User{}, errors.New("DATABASE_ERROR")
	} else if len(users) == 0 {
		return User{}, errors.New("MISSING_RECORD")
	}

	return users[0], nil
}

// SelectUsersRepository select all user data
func SelectUsersRepository() ([]User, error) {
	var users []User
	stmt := fmt.Sprintf("SELECT * FROM %s", userTable)
	err := mysqlDBHandler.Query(stmt, map[string]interface{}{}, &users)
	if err != nil {
		return users, errors.New("DATABASE_ERROR")
	} else if len(users) == 0 {
		return users, errors.New("MISSING_RECORD")
	}

	return users, nil
}

// UpdateUserRepository update user data
func UpdateUserRepository(data User) (User, error) {
	stmt := fmt.Sprintf("UPDATE %s SET first_name=:first_name,last_name=:last_name,contact_number=:contact_number WHERE id=:id", userTable)
	_, err := mysqlDBHandler.Execute(stmt, data)
	if err != nil {
		return User{}, errors.New("DATABASE_ERROR")
	}

	user := User{
		ID: data.ID,
	}

	return user, err
}

// DeleteUserRepository update user data
func DeleteUserRepository(data User) error {
	stmt := fmt.Sprintf("DELETE FROM %s WHERE id=:id", userTable)
	_, err := mysqlDBHandler.Execute(stmt, data)
	if err != nil {
		return errors.New("DATABASE_ERROR")
	}

	return err
}

// ============================== post repository ==============================

// InsertPostRepository insert a user data
func InsertPostRepository(data Post) (int64, error) {
	stmt := fmt.Sprintf("INSERT INTO %s (author_id, content) VALUES (:author_id, :content)", postTable)
	res, err := mysqlDBHandler.Execute(stmt, data)
	if err != nil {
		log.Println(err)
		if strings.Contains(err.Error(), "Duplicate entry") {
			return -1, errors.New("DUPLICATE_EMAIL")
		}

		return -1, errors.New("DATABASE_ERROR")
	}

	// get id
	id, err := res.LastInsertId()
	if err != nil {
		return -1, errors.New("DATABASE_ERROR")
	}

	return id, nil
}

// SelectPostByIDRepository select post data by id
func SelectPostByIDRepository(ID int64) (Post, error) {
	var posts []Post

	stmt := fmt.Sprintf("SELECT * FROM %s WHERE id=:id", postTable)
	err := mysqlDBHandler.Query(stmt, map[string]interface{}{
		"id": ID,
	}, &posts)
	if err != nil {
		return Post{}, errors.New("DATABASE_ERROR")
	} else if len(posts) == 0 {
		return Post{}, errors.New("MISSING_RECORD")
	}

	return posts[0], nil
}

// SelectPostsRepository select all user data
func SelectPostsRepository() ([]Post, error) {
	var posts []Post
	stmt := fmt.Sprintf("SELECT * FROM %s", postTable)
	err := mysqlDBHandler.Query(stmt, map[string]interface{}{}, &posts)
	if err != nil {
		return posts, errors.New("DATABASE_ERROR")
	} else if len(posts) == 0 {
		return posts, errors.New("MISSING_RECORD")
	}

	return posts, nil
}

// UpdatePostRepository update user data
func UpdatePostRepository(data Post) (Post, error) {
	stmt := fmt.Sprintf("UPDATE %s SET author_id=:author_id, content=:content WHERE id=:id", postTable)
	_, err := mysqlDBHandler.Execute(stmt, data)
	if err != nil {
		return Post{}, errors.New("DATABASE_ERROR")
	}

	post := Post{
		ID:       data.ID,
		AuthorID: data.AuthorID,
	}

	return post, err
}

// DeletePostRepository update user data
func DeletePostRepository(data Post) error {
	stmt := fmt.Sprintf("DELETE FROM %s WHERE id=:id", postTable)
	_, err := mysqlDBHandler.Execute(stmt, data)
	if err != nil {
		return errors.New("DATABASE_ERROR")
	}

	return err
}

// ============================== comment repository ==============================

// InsertCommentRepository insert a user data
func InsertCommentRepository(data Comment) (int64, error) {
	stmt := fmt.Sprintf("INSERT INTO %s (post_id, author_id, content) VALUES (:post_id, :author_id, :content)", commentTable)
	res, err := mysqlDBHandler.Execute(stmt, data)
	if err != nil {
		log.Println(err)
		if strings.Contains(err.Error(), "Duplicate entry") {
			return -1, errors.New("DUPLICATE_EMAIL")
		}

		return -1, errors.New("DATABASE_ERROR")
	}

	// get id
	id, err := res.LastInsertId()
	if err != nil {
		return -1, errors.New("DATABASE_ERROR")
	}

	return id, nil
}

// SelectCommentByIDRepository select post data by id
func SelectCommentByIDRepository(ID int64) (Comment, error) {
	var comments []Comment

	stmt := fmt.Sprintf("SELECT * FROM %s WHERE id=:id", commentTable)
	err := mysqlDBHandler.Query(stmt, map[string]interface{}{
		"id": ID,
	}, &comments)
	if err != nil {
		return Comment{}, errors.New("DATABASE_ERROR")
	} else if len(comments) == 0 {
		return Comment{}, errors.New("MISSING_RECORD")
	}

	return comments[0], nil
}

// SelectCommentsRepository select all user data
func SelectCommentsRepository() ([]Comment, error) {
	var comments []Comment
	stmt := fmt.Sprintf("SELECT * FROM %s", commentTable)
	err := mysqlDBHandler.Query(stmt, map[string]interface{}{}, &comments)
	if err != nil {
		return comments, errors.New("DATABASE_ERROR")
	} else if len(comments) == 0 {
		return comments, errors.New("MISSING_RECORD")
	}

	return comments, nil
}

// SelectCommentsByPostIDRepository select all user data
func SelectCommentsByPostIDRepository(postID int64) ([]Comment, error) {
	var comments []Comment
	stmt := fmt.Sprintf("SELECT * FROM %s WHERE post_id=:post_id", commentTable)
	err := mysqlDBHandler.Query(stmt, map[string]interface{}{
		"post_id": postID,
	}, &comments)
	if err != nil {
		return comments, errors.New("DATABASE_ERROR")
	} else if len(comments) == 0 {
		return comments, errors.New("MISSING_RECORD")
	}

	return comments, nil
}

// UpdateCommentRepository update user data
func UpdateCommentRepository(data Comment) (Comment, error) {
	stmt := fmt.Sprintf("UPDATE %s SET content=:content WHERE id=:id", commentTable)
	_, err := mysqlDBHandler.Execute(stmt, data)
	if err != nil {
		return Comment{}, errors.New("DATABASE_ERROR")
	}

	comment := Comment{
		ID:       data.ID,
		AuthorID: data.AuthorID,
	}

	return comment, err
}

// DeleteCommentRepository update user data
func DeleteCommentRepository(data Comment) error {
	stmt := fmt.Sprintf("DELETE FROM %s WHERE id=:id", commentTable)
	_, err := mysqlDBHandler.Execute(stmt, data)
	if err != nil {
		return errors.New("DATABASE_ERROR")
	}

	return err
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
