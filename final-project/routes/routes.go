package routes

import (
	photoControllers "fgd-golang/final-project/controllers/photo"
	"fgd-golang/final-project/middlewares"
	userControllers "fgd-golang/sesi-10/go-jwt/controllers/user"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", userControllers.UserRegister)
		userRouter.POST("/login", userControllers.UserLogin)

		userRouter.Use(middlewares.Authentication())
		userRouter.PUT("/:userId", middlewares.UserAuthorization(), userControllers.UpdateUser)
		userRouter.DELETE("/:userId", middlewares.UserAuthorization(), userControllers.DeleteUser)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/", photoControllers.CreatePhoto)
		photoRouter.GET("/", photoControllers.GetAllPhotos)

		photoRouter.PUT("/:photoId", photoControllers.UpdatePhoto)
		photoRouter.DELETE("/:photoId", photoControllers.DeletePhoto)
	}

	return r
}
