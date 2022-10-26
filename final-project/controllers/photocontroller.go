package controllers

import (
	"fgd-golang/final-project/helpers"
	"fgd-golang/final-project/models"
	"fgd-golang/final-project/repository"
	"log"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// var (
// 	photoJSON = "application/json"
// )

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

func UpdatePhoto(c *gin.Context) {
	db := repository.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}

	file, err := c.FormFile("photo_url")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(file.Filename)

	err = c.SaveUploadedFile(file, "assets/"+file.Filename)
	if err != nil {
		log.Fatal(err)
	}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID
	Photo.ID = uint(photoId)
	Photo.Photo_url = file.Filename

	err = db.Model(&Photo).Where("id = ?", photoId).Find(&Photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	if userID != Photo.UserID {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Bad Request",
			"message": "Not Found Data Id",
		})
		return
	}

	err = db.Model(&Photo).Where("id = ?", photoId).Updates(models.Photo{Title: Photo.Title, Caption: Photo.Caption, Photo_url: Photo.Photo_url}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Photo)
}

func DeletePhoto(c *gin.Context) {
	db := repository.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID
	Photo.ID = uint(photoId)

	err := db.Model(&Photo).Where("id = ?", photoId).Find(&Photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	if userID != Photo.UserID {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Bad Request",
			"message": "Not Found Data Id",
		})
		return
	}

	err = db.Model(&Photo).Where("id = ?", photoId).Delete(&Photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your Photo has been successfully deleted",
	})
}
