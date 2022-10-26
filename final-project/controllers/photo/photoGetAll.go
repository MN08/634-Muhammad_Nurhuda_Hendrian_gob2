package photoControllers

import (
	"fgd-golang/final-project/helpers"
	"fgd-golang/final-project/models"
	"fgd-golang/final-project/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllPhotos(c *gin.Context) {
	db := repository.GetDB()
	// userData := c.MustGet("userData").(jwt.MapClaims)
	Photo := []models.Photo{}
	contentType := helpers.GetContentType(c)
	// userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	// Photo.UserID = userID

	err := db.Find(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Photo)

}
