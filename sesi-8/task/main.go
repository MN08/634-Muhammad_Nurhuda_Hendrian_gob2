package main

import (
	"fgd-golang/sesi-8/task/config"
	"fgd-golang/sesi-8/task/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()
	router.GET("/person/:id", inDB.GetPerson)
	router.GET("/persons", inDB.GetPerson)
	router.POST("/person", inDB.CreatePerson)
	router.PUT("/person", inDB.UpdatePerson)
	router.DELETE("/person/:id", inDB.DeletePerson)
	router.Run(":3000")

}
