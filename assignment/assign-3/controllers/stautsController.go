package controllers

import (
	"fgd-golang/assignment/assign-3/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Data struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func (idb *InDB) CreateData(c *gin.Context) {
	var (
		stat   models.Status
		result gin.H
	)

	water := c.PostForm("water")
	wind := c.PostForm("wind")

	intWater, _ := strconv.Atoi(water)
	intWind, _ := strconv.Atoi(wind)

	stat.Water = intWater
	stat.Wind = intWind

	dataStat := Data{
		Water: stat.Water,
		Wind:  stat.Wind,
	}

	err := idb.DB.Create(&stat).Error

	if err != nil {
		result = gin.H{
			"result": "Data cannot be inserted",
		}
		c.JSON(http.StatusNoContent, result)
	} else {
		result = gin.H{
			"status": dataStat,
		}
		c.JSON(http.StatusOK, result)
	}
}

func (idb *InDB) GetLatestData(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")

	var (
		stat    models.Status
		allStat []models.Status
		result  gin.H
	)

	idb.DB.Find(&allStat).Take(&allStat).Last(&allStat)

	if len(allStat) <= 0 {
		result = gin.H{
			"message": "Data not found!",
		}
		c.JSON(http.StatusNotFound, result)
	} else {
		idb.DB.Find(&stat).Take(&stat).Last(&stat)

		var waterStatus, windStatus string

		dataStat := Data{
			Water: stat.Water,
			Wind:  stat.Wind,
		}

		if dataStat.Water <= 5 {
			waterStatus = "Aman"
		} else if dataStat.Water >= 6 && dataStat.Water <= 8 {
			waterStatus = "Siaga"
		} else if dataStat.Water > 8 {
			waterStatus = "Bahaya"
		}

		if dataStat.Wind <= 6 {
			windStatus = "Aman"
		} else if dataStat.Wind >= 7 && dataStat.Wind <= 15 {
			windStatus = "Siaga"
		} else if dataStat.Wind > 15 {
			windStatus = "Bahaya"
		}

		result = gin.H{
			"status":    dataStat,
			"waterStat": waterStatus,
			"windStat":  windStatus,
		}

		c.JSON(http.StatusOK, result)

	}

}
