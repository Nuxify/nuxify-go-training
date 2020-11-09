package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// HTTPResponseVM base http viewmodel for http rest responses
type HTTPResponseVM struct {
	Status    int         `json:"-"`
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	ErrorCode interface{} `json:"errorCode,omitempty"`
	Data      interface{} `json:"data"`
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
	})

	fmt.Println("Server is listening on " + port)
	log.Fatal(http.ListenAndServe(port, router))
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

// JSON converts http responsewriter to json
func (response *HTTPResponseVM) JSON(w http.ResponseWriter) {
	if response.Data == nil {
		response.Data = map[string]interface{}{}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Status)
	_ = json.NewEncoder(w).Encode(response)
}
