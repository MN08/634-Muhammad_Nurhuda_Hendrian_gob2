package route

import (
	"fgd-golang/task/satu/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/hello", controllers.CreateBio)
	router.GET("/hello", controllers.GetBio)

	return router
}
