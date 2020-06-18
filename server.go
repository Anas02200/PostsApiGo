package main

import (
	"GoApi/controller"
	"GoApi/http"
	"fmt"
	"log"
	"net/http"
)

var (
	postController controller.PostController = controller.NewPostController()
	httpRouter     router.Router             = router.NewMuxRouter()
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
