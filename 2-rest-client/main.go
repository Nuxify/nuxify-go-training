package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type values struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int    `json:"userId"`
}

func main() {
	CreatePost()
	GetAllPosts()
	GetPostByID()
	DeleteByID()
	UpdatePost()
}

// CreatePost create a post
func CreatePost() {
	book := values{"GarzAlma", "bar", 1}

	jsonValue, err := json.Marshal(book)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var response map[string]interface{}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		panic(err)
	}

	fmt.Println("POST REQUEST:", response)

}

// DeleteByID delete id
func DeleteByID() {
	book := values{"GarzAlma", "bar", 1}

	jsonValue, err := json.Marshal(book)
	req, err := http.NewRequest(http.MethodDelete, "https://jsonplaceholder.typicode.com/posts/1", bytes.NewBuffer(jsonValue))
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

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {

		panic(err)
	}

	fmt.Println(string(body))
}

// GetPostByID get post by id
func GetPostByID() {

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println("API Response as String:\n" + bodyString)

	// Convert response body to map interface
	var response map[string]interface{}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		panic(err)
	}

	fmt.Println("POST REQUEST:", response)
}

// UpdatePost update post by id
func UpdatePost() {
	book := values{"GarzAlma", "bar", 1}

	// initialize http client
	client := &http.Client{}

	// marshal User to json
	jsonValue, err := json.Marshal(book)
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
}
