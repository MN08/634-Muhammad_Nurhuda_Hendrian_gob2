package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Car struct {
	CarId string `json:"car_id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Price int    `json:"price"`
}

var carDatas = []Car{}

func CreateCar(ctx *gin.Context) {
	var newCar Car
	if err := ctx.ShouldBindJSON(&newCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	newCar.CarId = fmt.Sprintf("c%d", len(carDatas)+1)
	carDatas = append(carDatas, newCar)

	ctx.JSON(http.StatusCreated, gin.H{
		"cars": newCar,
	})
}

func UpdateCar(ctx *gin.Context) {
	carId := ctx.Param("carId")
	condition := false
	var updateCar Car

	if err := ctx.ShouldBindJSON(&updateCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, car := range carDatas {
		if carId == car.CarId {
			condition = true
			carDatas[i] = updateCar
			carDatas[i].CarId = carId
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error status":  "Data Not Found",
			"error meesage": fmt.Sprintf("car with id %v not found", carId),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("car with id %v has been successfully updated", carId),
	})
}

func GetCar(ctx *gin.Context) {
	carId := ctx.Param("carId")
	condition := false
	var carData Car

	for i, car := range carDatas {
		if carId == car.CarId {
			condition = true
			carData = carDatas[i]
			break
		}
	}
	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error status":  "Data Not Found",
			"error meesage": fmt.Sprintf("car with id %v Not found", carId),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"car": carData,
	})
}

func GetAllCar(ctx *gin.Context) {
	var carData Car

	for i := range carDatas {
		carData = carDatas[i]

		ctx.JSON(http.StatusOK, gin.H{
			"car": carData,
		})
	}
}

func DeleteCar(ctx *gin.Context) {
	carId := ctx.Param("carId")
	condition := false
	var carIndex int

	for i, car := range carDatas {
		if carId == car.CarId {
			condition = true
			carIndex = i
			break
		}
	}
	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error status":  "Data Not Found",
			"error meesage": fmt.Sprintf("car with id %v Not found", carId),
		})
		return
	}

	copy(carDatas[carIndex:], carDatas[carIndex+1:])
	carDatas[len(carDatas)-1] = Car{}
	carDatas = carDatas[:len(carDatas)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("car with id %v has been successfully deleted", carId),
	})
}
