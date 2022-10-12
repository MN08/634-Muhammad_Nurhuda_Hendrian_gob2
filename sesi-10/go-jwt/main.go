package main

import (
	"fgd-golang/sesi-10/go-jwt/database"
	"fgd-golang/sesi-10/go-jwt/routers"
)

func main() {
	database.StartDb()
	r := routers.StartApp()
	r.Run(":8000")
}
