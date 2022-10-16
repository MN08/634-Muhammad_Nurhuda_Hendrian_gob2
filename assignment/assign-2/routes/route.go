package route

import (
	orderControllers "fgd-golang/assignment/assign-2/controllers/order"
	"fgd-golang/assignment/assign-2/database"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {

	db := database.StartDb()
	inDB := &orderControllers.InDB{DB: db}
	router := gin.Default()

	router.POST("/orders", inDB.CreateOrder)
	router.GET("/orders", inDB.GetOrder)
	router.PUT("/orders/:orderId", inDB.UpdateOrder)
	router.DELETE("/orders/:orderId", inDB.DeleteOrder)
	router.Run(":3000")

	return router
}
