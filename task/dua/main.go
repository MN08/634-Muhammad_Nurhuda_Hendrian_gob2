package main

import (
	controllers "fgd-golang/task/dua/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	api := gin.Default()
	api.POST("/hello", controllers.CreateBio)
	api.GET("/hello", controllers.Middleware1())
	api.GET("/hellos", controllers.Middleware2())
	api.GET("/sesisatu", controllers.Satu())
	api.Run(":8080")
}
