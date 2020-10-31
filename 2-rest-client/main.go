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

	var response map[string]interface{}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		panic(err)
	}

	fmt.Println("POST REQUEST:", response)

}

// DeleteByID delete id
func DeleteByID() {
	values := values{"GarzAlma", "bar", 1}
	jsonReq, err := json.Marshal(values)
	req, err := http.NewRequest(http.MethodDelete, "https://jsonplaceholder.typicode.com/posts/1", bytes.NewBuffer(jsonReq))
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

// UpdatePost update post by id
func UpdatePost() {

}
