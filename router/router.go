package router

import (
	"my-gram/controller"
	"my-gram/middleware"

	_ "my-gram/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MyGram API
// @version 1.0
// @description this is sample services for MyGram
// @termsOfService http://swagger.io/terms/
// @host localholst:8080
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /
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
		photoRouter.PUT("/:photoId", middleware.PhotoOwnerAuth(), controller.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middleware.PhotoOwnerAuth(), controller.DeleteProduct)
		photoRouter.GET("/", controller.GetAllPhoto)
	}

	sosmedRouter := r.Group("/sosmed")
	{
		sosmedRouter.Use(middleware.Authentication())
		sosmedRouter.POST("/", controller.CreateSosmed)
		sosmedRouter.GET("/:sosmedId", controller.GetSosmed)
		sosmedRouter.PUT("/:sosmedId", middleware.SosmedOwnerAuth(), controller.UpdateSosmed)
		sosmedRouter.DELETE("/:sosmedId", middleware.SosmedOwnerAuth(), controller.DeleteSosmed)
		sosmedRouter.GET("/", controller.GetAllSosmed)
	}

	commentRouter := r.Group("/comment")
	{
		commentRouter.Use(middleware.Authentication())
		commentRouter.POST("/", controller.CreateComment)
		commentRouter.GET("/:commentId", controller.GetComment)
		commentRouter.PUT("/:commentId", middleware.CommentOwnerAuth(), controller.UpdateComment)
		commentRouter.DELETE("/:commentId", middleware.CommentOwnerAuth(), controller.DeleteComment)
		commentRouter.GET("/", controller.GetAllCcomment)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}
