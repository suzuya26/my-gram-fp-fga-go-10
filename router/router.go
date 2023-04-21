package router

import (
	"my-gram/controller"
	"my-gram/middleware"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controller.UserRegister)
		userRouter.POST("/login", controller.UserLogin)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middleware.Authentication())
		photoRouter.POST("/", controller.CreatePhoto)
		photoRouter.GET("/:photoId", controller.GetPhoto)
		photoRouter.PUT("/:photoId", controller.UpdatePhoto)
		photoRouter.DELETE("/:photoId", controller.DeleteProduct)
		photoRouter.GET("/", controller.GetAllPhoto)
	}

	sosmedRouter := r.Group("/sosmed")
	{
		sosmedRouter.Use(middleware.Authentication())
		sosmedRouter.POST("/", controller.CreateSosmed)
		sosmedRouter.GET("/:sosmedId", controller.GetSosmed)
		sosmedRouter.PUT("/:sosmedId", controller.UpdateSosmed)
		sosmedRouter.DELETE("/:sosmedId", controller.DeleteSosmed)
		sosmedRouter.GET("/", controller.GetAllSosmed)
	}

	commentRouter := r.Group("/comment")
	{
		commentRouter.Use(middleware.Authentication())
		commentRouter.POST("/", controller.CreateComment)
		commentRouter.GET("/:commentId", controller.GetComment)
		commentRouter.PUT("/:commentId", controller.UpdateComment)
		commentRouter.DELETE("/:commentId", controller.DeleteComment)
		commentRouter.GET("/", controller.GetAllCcomment)
	}

	return r
}
