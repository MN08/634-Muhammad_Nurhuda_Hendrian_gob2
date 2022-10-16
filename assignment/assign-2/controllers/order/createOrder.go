package orderControllers

import (
	"fgd-golang/assignment/assign-2/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Data struct {
	ItemCode    string `json:"itemCode"`
	Description string `json:"quantity"`
	Quantity    int    `json:"quantity"`
}

func (idb *InDB) CreateOrder(c *gin.Context) {

	var (
		order  models.Order
		item   models.Item
		result gin.H
	)

	customerName := c.PostForm("customerName")
	itemCode := c.PostForm("itemCode")
	description := c.PostForm("description")
	sQuantity := c.PostForm("quantity")

	quantity, _ := strconv.Atoi(sQuantity)

	order.OrderedAt = time.Now()
	order.CustomerName = customerName
	item.Code = itemCode
	item.Description = description
	item.Quantity = quantity

	var data []Data
	dataStruct := Data{
		ItemCode:    item.Code,
		Description: item.Description,
		Quantity:    quantity,
	}

	data = append(data, dataStruct)

	err := idb.DB.Create(&order).Error
	if err != nil {
		result = gin.H{
			"result": "Cannot input new data",
		}
		c.JSON(http.StatusNoContent, result)
	} else {
		item.OrderID = order.ID

		err = idb.DB.Create(&item).Error
		if err != nil {
			result = gin.H{
				"result": "Cannot input new data",
			}
			c.JSON(http.StatusNoContent, result)
		} else {

			c.JSON(http.StatusOK, gin.H{
				"orderId":      order.ID,
				"orderedAt":    order.OrderedAt,
				"customerName": order.CustomerName,
				"items":        data,
			})
		}
	}
}
