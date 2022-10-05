package main

import (
	productControllers "fgd-golang/sesi-7/gorm/controllers/products"
	userControllers "fgd-golang/sesi-7/gorm/controllers/users"
	"fgd-golang/sesi-7/gorm/databases"
	"fgd-golang/sesi-7/gorm/models"
	"fmt"
)

func main() {
	databases.StartDb()
	//create User to Db
	userControllers.CreateUser("andre@mjd.online")
	//get single user from database
	userControllers.GetUserById(1)
	//update User data
	userControllers.UpdateUser(1, "andrea@mjd.online")

	//create products data
	productControllers.CreateProduct(1, "undiscribe", "unknown")

	//get user with products
	getUserWithProduct()

}

func getUserWithProduct() {
	db := databases.GetDB()

	users := models.User{}
	err := db.Preload("Products").Find(&users).Error

	if err != nil {
		fmt.Println("Error getting user with Products", err.Error())
		return
	}

	fmt.Println("user with product")
	fmt.Printf("%+v\n", users)
}
