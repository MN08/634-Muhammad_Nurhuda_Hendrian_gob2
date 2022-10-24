package routers

import (
	productControllers "fgd-golang/sesi-10/go-jwt/controllers/product"
	userControllers "fgd-golang/sesi-10/go-jwt/controllers/user"
	"fgd-golang/sesi-10/go-jwt/middleware"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", userControllers.UserRegister)
		userRouter.POST("/login", userControllers.UserLogin)
	}

	productRouter := r.Group("/products")
	{
		productRouter.Use(middleware.Authentication())
		productRouter.POST("/", productControllers.CreateProduct)

		productRouter.PUT("/:productId", middleware.ProductAuthorization(), productControllers.UpdateProduct)
	}

	return r
}
