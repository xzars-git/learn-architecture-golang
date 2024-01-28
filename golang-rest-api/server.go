package main

import (
	"fmt"
	"learn-architecture-golang/golang-rest-api/golang-rest-api/controller"
	router "learn-architecture-golang/golang-rest-api/golang-rest-api/http"
	"learn-architecture-golang/golang-rest-api/golang-rest-api/repository"
	"learn-architecture-golang/golang-rest-api/golang-rest-api/service"
	"net/http"
)

var (
	postRepository repository.PostRepository = repository.NewFirestoreRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter     router.Router             = router.NewChiRouter()
)

func main() {
	const port string = ":8000"
	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})
	httpRouter.GET("/posts", postController.GetPost)
	httpRouter.POST("/posts", postController.AddPost)
	httpRouter.SERVE(port)

}
