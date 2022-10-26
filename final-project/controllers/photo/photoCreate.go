package photoControllers

import (
	"fgd-golang/final-project/helpers"
	"fgd-golang/final-project/models"
	"fgd-golang/final-project/repository"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func CreatePhoto(c *gin.Context) {
	db := repository.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	file, err := c.FormFile("photo_url")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(file.Filename)

	err = c.SaveUploadedFile(file, "assets/"+file.Filename)
	if err != nil {
		log.Fatal(err)
	}

	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID
	Photo.Photo_url = file.Filename
	err = db.Debug().Create(&Photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, Photo)
	// c.JSON(http.StatusCreated, gin.H{
	// 	"id":               Photo.ID,
	// 	"title":            Photo.Title,
	// 	"social_media_url": Photo.Caption,
	// 	"photo_url":        file.Filename,
	// 	"user_id":          userID,
	// 	"created_at":       Photo.CreatedAt,
	// })
}
