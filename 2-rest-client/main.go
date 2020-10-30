package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	CreatePost()
}

// CreatePost create a post
func CreatePost() {
	values := map[string]interface{}{
		"title":  "GarzAlma",
		"body":   "bar",
		"userId": 1,
	}

	jsonValue, err := json.Marshal(values)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Status Code", resp.StatusCode)
	fmt.Println("Body", resp.Body)

	var response map[string]interface{}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		panic(err)
	}

	fmt.Println("POST REQUEST:", response)

}

// DeleteByID delete id
func DeleteByID() {

}

// GetPostByID get post by id
func GetPostByID() {

}

// GetAllPosts get all posts
func GetAllPosts() {

}

// UpdatePost update post by id
func UpdatePost() {

}
