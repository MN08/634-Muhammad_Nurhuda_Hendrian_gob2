package userControllers

import (
	"fgd-golang/final-project/helpers"
	"fgd-golang/final-project/models"
	"fgd-golang/final-project/repository"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func UpdateUser(c *gin.Context) {
	db := repository.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	userId := uint(userData["id"].(float64))
	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}
	User.ID = userId

	err := db.Model(&User).Where("id = ?", userId).Updates(models.User{Email: User.Email, Username: User.Username, Age: User.Age}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":       User.ID,
		"email":    User.Email,
		"username": User.Username,
		"age":      User.Age,
	})

}
