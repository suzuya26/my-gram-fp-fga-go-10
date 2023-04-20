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
		photoRouter.POST("/", controller.Createphoto)
		photoRouter.GET("/:photoId", controller.GetPhoto)
		photoRouter.PUT("/:photoId", controller.UpdatePhoto)
		photoRouter.DELETE("/:photoId", controller.DeleteProduct)
		photoRouter.GET("/", controller.GetAllPhoto)
	}

	return r
}
