package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	port := ":8080"

	fmt.Println("Starting Server....")

	router := chi.NewRouter()
	router.Get("/api/posts", getAllHandler)
	router.Get("/api/posts/{id}", getByIDHandler)
	router.Post("/api/posts", postHandler)
	router.Patch("/api/posts/{id}", patchHandler)
	router.Delete("/api/posts/{id}", delhHandler)
	fmt.Println("Server is listening on " + port)
	log.Fatal(http.ListenAndServe(port, router))

}

func delhHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Your request has been deleted!")
}

func getAllHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("You got me a GET Request!")
}

func getByIDHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("You got me a GET Request!")
}

func patchHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Your request has been updated!")
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("You just sent me a post request!")
}
