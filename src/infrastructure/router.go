package router

import (
	"crud-go/src/controller"
	"crud-go/src/repository"
	"crud-go/src/usecase"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	postRepo := repository.NewPostRepositoryDB()
	postUsecase := usecase.NewPostUsecase(postRepo)

	healthController := &controller.HealthController{}
	postController := SetupPostController(router, postUsecase)

	router.GET("/health", healthController.GetHealth)
	router.GET("/posts", postController.GetPosts)
	router.GET("/posts/:id", postController.GetPostById)
	router.POST("/posts", postController.CreatePost)

	return router
}

func SetupPostController(router *gin.Engine, pu usecase.PostUsecase) *controller.PostController {
	return &controller.PostController{
		PostUsecase: pu,
	}
}
