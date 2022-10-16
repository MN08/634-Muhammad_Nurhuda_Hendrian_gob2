package orderControllers

import (
	"fgd-golang/assignment/assign-2/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetOrder(c *gin.Context) {

	var (
		orders []models.Order
		result gin.H
	)

	idb.DB.Find(&orders)

	if len(orders) <= 0 {
		result = gin.H{
			"result":  nil,
			"count":   0,
			"message": "Data not found",
		}
		c.JSON(http.StatusNotFound, result)
	} else {

		result = gin.H{
			"result": orders,
			"count":  len(orders),
		}
		c.JSON(http.StatusOK, result)
	}
}
