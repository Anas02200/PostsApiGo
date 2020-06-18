package main

import (
	"GoApi/controller"
	"GoApi/http"
	"GoApi/repository"
	"GoApi/services"
	"fmt"
	"log"
	"net/http"
)

var (
	postRepository repository.PostRepository = repository.NewFirestoreRepository()
	PostService    services.PostService      = services.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(PostService)
	httpRouter     router.Router             = router.NewChiRouter()
)

func main() {

	const port string = ":3500"
	httpRouter.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprint(resp, "up")
	})
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)
	log.Println("listenning on port : ", port)
	httpRouter.SERVE(port)

}
