package userControllers

import (
	"fgd-golang/sesi-7/gorm/databases"
	"fgd-golang/sesi-7/gorm/models"
	"fmt"
)

func CreateUser(email string) {
	db := databases.GetDB()
	fmt.Println(db)
	User := models.User{
		Email: email,
	}

	err := db.Create(&User).Error
	if err != nil {
		fmt.Println("Error creating user data:", err)
		return
	}

	fmt.Println("New user data created", User)

}
