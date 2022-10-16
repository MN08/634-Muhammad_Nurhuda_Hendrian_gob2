package orderControllers

import (
	"fgd-golang/assignment/assign-2/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (idb *InDB) DeleteOrder(c *gin.Context) {

	var (
		order  models.Order
		item   models.Item
		result gin.H
	)

	orderId := c.Param("orderId")

	err := idb.DB.Where("id = ?", orderId).First(&order).Error
	if err != nil {

		result = gin.H{
			"result": "Data not found",
		}
		c.JSON(http.StatusNotFound, result)
	} else {

		err = idb.DB.Where("orderId = ?", orderId).First(&item).Error
		if err != nil {

			result = gin.H{
				"result": "Data not found",
			}
			c.JSON(http.StatusNotFound, result)
		} else {

			err = idb.DB.Delete(&item).Error
			if err != nil {

				result = gin.H{
					"result": "Failed delete data",
				}
				c.JSON(http.StatusBadRequest, result)
			} else {

				err = idb.DB.Delete(&order).Error
				if err != nil {

					result = gin.H{
						"result": "Failed delete data",
					}
					c.JSON(http.StatusBadRequest, result)
				} else {

					result = gin.H{
						"result": "Data has been deleted",
						"data":   order,
					}
					c.JSON(http.StatusOK, result)
				}
			}
		}
	}
}
