package main

import (
	"fgd-golang/final-project/repository"
	"fgd-golang/final-project/routes"
)

func main() {
	repository.StartDb()
	r := routes.StartApp()
	r.Run(":8080")
}
