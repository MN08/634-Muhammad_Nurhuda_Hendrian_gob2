package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Bio struct {
	ID       string `json:"id"`
	FullName string `json:"fullName"`
	RegID    string `json:"regID"`
}

var bioDatas = []Bio{}

func CreateBio(ctx *gin.Context) {
	var newBio Bio
	if err := ctx.ShouldBindJSON(&newBio); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	newBio.ID = fmt.Sprintf("%d", len(bioDatas)+1)
	bioDatas = append(bioDatas, newBio)

	ctx.JSON(http.StatusCreated, gin.H{
		"bio": newBio,
	})
}

func GetBio(ctx *gin.Context) {
	var bioData Bio

	for i := range bioDatas {
		bioData = bioDatas[i]

		ctx.JSON(http.StatusOK, gin.H{
			"bio": bioData,
		})
	}
}
