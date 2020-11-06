package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	fmt.Println("Starting Server....")

	router := chi.NewRouter()
	router.Get("/api/getExample", getHandler)
	router.Post("/api/createExample", postHandler)
	router.Patch("/api/updateExample", patchHandler)
	router.Delete("/api/deleteExample", delhHandler)
	fmt.Println("Server is listening on port 8080....")
	log.Fatal(http.ListenAndServe(":8080", router))

}

func delhHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Your request has been deleted!")
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("You got me a GET Request!")
}

func patchHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Your request has been updated!")
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("You just sent me a post request!")
}
