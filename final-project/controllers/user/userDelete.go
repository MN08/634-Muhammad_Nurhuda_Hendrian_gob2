package userControllers

import (
	"fgd-golang/final-project/helpers"
	"fgd-golang/final-project/models"
	"fgd-golang/final-project/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// var (
// 	appJSON = "application/json"
// )

func DeleteUser(c *gin.Context) {
	db := repository.GetDB()
	contentType := helpers.GetContentType(c)
	User := models.User{}

	userId, _ := strconv.Atoi(c.Param("userId"))

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}
	User.ID = uint(userId)

	err := db.Where("id = ?", userId).Delete(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Your account has been sucessfully deleted",
	})
}
