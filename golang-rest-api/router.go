package main

import (
	"encoding/json"
	"learn-architecture-golang/golang-rest-api/golang-rest-api/entity"
	"learn-architecture-golang/golang-rest-api/golang-rest-api/repository"
	"math/rand"
	"net/http"
)

var (
	posts []entity.Post
)

var (
	repo repository.PostRepository = repository.NewPostRepository()
)

func getPost(response http.ResponseWriter, req *http.Request) {

	response.Header().Set("Content-Type", "application/json")
	posts, err := repo.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error getting the posts"}`))
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(posts)
}

func addPost(response http.ResponseWriter, req *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var post entity.Post
	err := json.NewDecoder(req.Body).Decode(&post)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error unmarshaling the request"}`))
		return
	}
	post.ID = rand.Int63()
	repo.Save(&post)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(post)
}
