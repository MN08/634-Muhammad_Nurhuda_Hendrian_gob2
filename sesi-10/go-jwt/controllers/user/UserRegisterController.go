package userControllers

import (
	"fgd-golang/sesi-10/go-jwt/database"
	"fgd-golang/sesi-10/go-jwt/helpers"
	"fgd-golang/sesi-10/go-jwt/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"ID":        User.ID,
		"email":     User.Email,
		"full_name": User.FullName,
	})

}
