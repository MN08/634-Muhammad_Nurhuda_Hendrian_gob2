package orderControllers

import (
	"fgd-golang/assignment/assign-2/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (idb *InDB) UpdateOrder(c *gin.Context) {

	id := c.Param("orderId")

	type UpdateData struct {
		LineItemId  int    `json:"lineItemId"`
		ItemCode    string `json:"itemCode"`
		Description string `json:"description"`
		Quantity    int    `json:"quantity"`
	}

	var (
		order    models.Order
		item     models.Item
		newOrder models.Order
		result   gin.H
	)

	err := idb.DB.Where("id = ?", id).First(&order).Error
	if err != nil {

		result = gin.H{
			"result": "Data not found",
		}
		c.JSON(http.StatusNotFound, result)
	} else {

		err = idb.DB.Where("order_id = ?", id).First(&item).Error
		if err != nil {

			result = gin.H{
				"result": "Data not found",
			}
			c.JSON(http.StatusNotFound, result)

		} else {

			orderid, _ := strconv.Atoi(id)
			customerName := c.PostForm("customerName")
			newOrder.ID = orderid
			newOrder.CustomerName = customerName
			newOrder.OrderedAt = order.OrderedAt

			var data []UpdateData
			dataStruct := UpdateData{
				LineItemId:  item.ID,
				ItemCode:    item.Code,
				Description: item.Description,
				Quantity:    item.Quantity,
			}

			data = append(data, dataStruct)

			err = idb.DB.Model(&order).Updates(newOrder).Error
			if err != nil {

				result = gin.H{
					"result": "Update Failed",
				}
				c.JSON(http.StatusNoContent, result)
			} else {

				result = gin.H{
					"result":        "Data updated!",
					"orderId":       order.ID,
					"orderedAt":     order.OrderedAt,
					"customerName": order.CustomerName,
					"data":          data,
				}
				c.JSON(http.StatusOK, result)
			}
		}
	}

}
