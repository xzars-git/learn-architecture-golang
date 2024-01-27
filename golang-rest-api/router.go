package main

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

var (
	posts []Post
)

func init() {
	posts = []Post{
		{Id: 1, Title: "Title1", Text: "Text1"},
	}
}

func getPost(response http.ResponseWriter, req *http.Request) {

	response.Header().Set("Content-Type", "application/json")
	result, err := json.Marshal(posts)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error marshaling the posts array"}`))
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Write(result)
}

func addPost(response http.ResponseWriter, req *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var post Post
	err := json.NewDecoder(req.Body).Decode(&post)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error unmarshaling the request"}`))
		return
	}
	post.Id = len(posts) + 1
	posts = append(posts, post)
	response.WriteHeader(http.StatusOK)
	result, err := json.Marshal(post)
	response.Write(result)
}
