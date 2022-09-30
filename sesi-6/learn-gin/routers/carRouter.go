package routers

import (
	"learn-gin/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/cars", controllers.CreateCar)
	router.PUT("/cars/:carId", controllers.UpdateCar)
	router.GET("/cars/:carId", controllers.GetCar)
	router.GET("/cars", controllers.GetAllCar)
	router.DELETE("/cars/:carId", controllers.DeleteCar)

	return router
}
