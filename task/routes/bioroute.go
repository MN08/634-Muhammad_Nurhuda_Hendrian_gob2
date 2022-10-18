package route

import (
	"fgd-golang/task/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	route := gin.Default()

	route.POST("/hello", controllers.CreateBio)
	route.GET("/hello/:ID", controllers.GetBio)
	// router.PUT("/cars/:carId", controllers.UpdateCar)
	// router.GET("/cars", controllers.GetAllCar)
	// router.DELETE("/cars/:carId", controllers.DeleteCar)

	return route
}
