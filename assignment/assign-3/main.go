package main

import (
	"fgd-golang/assignment/assign-3/controllers"
	"fgd-golang/assignment/assign-3/database"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var PORT = ":8080"
	db := database.StartDb()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.POST("/stat/", inDB.CreateData)
	router.GET("/stat/latest", inDB.GetLatestData)

	router.Run(PORT)
}
