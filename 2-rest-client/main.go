package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
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

var (
	client *http.Client = &http.Client{Timeout: 3 * time.Second}
)

func main() {
	// CreatePost()
	// GetAllPosts()
	// GetPostByID()
	// DeleteByID()
	UpdatePost()
}

// CreatePost create a post
func CreatePost() {
	var result response
	payload := createRequestPayload{
		Title:  "GarzAlma",
		Body:   "bar",
		UserID: 1,
	}
	err := HTTPPost("https://jsonplaceholder.typicode.com/posts", &payload, &result)
	if err != nil {
		panic(err)
	}
	fmt.Println("POST REQUEST:", result.Title)
}

// DeleteByID delete id
func DeleteByID() {
	var result response
	err := HTTPDelete("https://jsonplaceholder.typicode.com/posts/1", &result)
	if err != nil {
		panic(err)
	}
	fmt.Println("POST REQUEST:", result)
}

// GetAllPosts get all posts
func GetAllPosts() {
	var results []response
	err := HTTPGet("https://jsonplaceholder.typicode.com/posts", &results)
	if err != nil {
		panic(err)
	}
	fmt.Println("GET ALL POSTS", results)
}

// GetPostByID get post by id
func GetPostByID() {
	var result response
	err := HTTPGet("https://jsonplaceholder.typicode.com/posts/1", &result)
	if err != nil {
		panic(err)
	}
	fmt.Println("POST REQUEST:", result)
}

// UpdatePost update post by id
func UpdatePost() {
	var result response
	payload := updateRequestPayload{Title: "GarzAlma update into GarzClang"}

	err := HTTPPatch("https://jsonplaceholder.typicode.com/posts/1", &payload, &result)
	if err != nil {
		panic(err)
	}
	fmt.Println("POST REQUEST:", result.Title)
}

// HTTPGet is a http get request helper
func HTTPGet(url string, response interface{}) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return err
	}
	return nil
}

// HTTPDelete is a http get request helper
func HTTPDelete(url string, response interface{}) error {
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return err
	}
	return nil
}

// HTTPPost is a http post request helper
func HTTPPost(url string, payload, response interface{}) error {
	// convert to reader
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(payload)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, url, buf)
	if err != nil {
		return err
	}
	// add custom headers
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return err
	}
	return nil
}

// HTTPPatch is a http post request helper
func HTTPPatch(url string, payload, response interface{}) error {
	// convert to reader
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(payload)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPatch, url, buf)
	if err != nil {
		return err
	}
	// add custom headers
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return err
	}
	return nil
}
