package controllers

import (
	"fgd-golang/sesi-8/task/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (idb *InDB) DeletePerson(c *gin.Context) {
	var (
		person models.Person
		result gin.H
	)

	id := c.Param("id")
	err := idb.DB.First(&person, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	err = idb.DB.Delete(&person).Error
	if err != nil {
		result = gin.H{
			"result": "delete failed",
		}
	} else {
		result = gin.H{
			"result": "successfully deleted",
		}
	}
	c.JSON(http.StatusOK, result)
}
