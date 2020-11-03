package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type createRequestPayload struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int    `json:"userId"`
}
type response struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int    `json:"userId"`
}
type updateRequestPayload struct {
	Title string `json:"title"`
}

func main() {
	// CreatePost()
	// GetAllPosts()
	// GetPostByID()
	// DeleteByID()
	UpdatePost()
}

// CreatePost create a post
func CreatePost() {
	payload := createRequestPayload{
		Title:  "GarzAlma",
		Body:   "bar",
		UserID: 1,
	}
	jsonValue, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}
	resp, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var result response
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("POST REQUEST:", result.Title)
}

// DeleteByID delete id
func DeleteByID() {
	req, err := http.NewRequest(http.MethodDelete, "https://jsonplaceholder.typicode.com/posts/1", nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
}

// GetAllPosts get all posts
func GetAllPosts() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var results []response
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		panic(err)
	}
	fmt.Println("GET ALL POSTS", results)
}

// GetPostByID get post by id
func GetPostByID() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println("API Response as String:\n" + bodyString)
	var results []response
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		panic(err)
	}
	fmt.Println("POST REQUEST:", results)
}

// UpdatePost update post by id
func UpdatePost() {
	payload := updateRequestPayload{Title: "Update GarzAlma into GarzClang"}
	client := &http.Client{}
	// marshal User to json
	jsonValue, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	// set the HTTP method, url, and request body
	req, err := http.NewRequest(http.MethodPatch, "https://jsonplaceholder.typicode.com/posts/1", bytes.NewBuffer(jsonValue))
	if err != nil {
		panic(err)
	}
	// set the request header Content-Type for json
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println("API Response as String:\n" + bodyString)
	var results []response
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		panic(err)
	}
	fmt.Println("POST REQUEST:", results)
}
