package controller

import (
	"encoding/json"
	"learn-architecture-golang/golang-rest-api/golang-rest-api/entity"
	"learn-architecture-golang/golang-rest-api/golang-rest-api/errors"
	"learn-architecture-golang/golang-rest-api/golang-rest-api/service"
	"net/http"
)

type controller struct{}

var (
	postService service.PostService
)

type PostController interface {
	GetPost(response http.ResponseWriter, req *http.Request)
	AddPost(response http.ResponseWriter, req *http.Request)
}

func NewPostController(service service.PostService) PostController {
	postService = service
	return &controller{}
}

func (*controller) GetPost(response http.ResponseWriter, req *http.Request) {

	response.Header().Set("Content-Type", "application/json")
	posts, err := postService.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error getting the posts"})
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(posts)
}

func (*controller) AddPost(response http.ResponseWriter, req *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var post entity.Post
	err := json.NewDecoder(req.Body).Decode(&post)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error unmarshaling the request"})
		return
	}

	err1 := postService.Validate(&post)
	if err1 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: err1.Error()})
		return
	}

	result, err2 := postService.Create(&post)
	if err2 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error saving the post"})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}
