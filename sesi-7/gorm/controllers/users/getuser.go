package userControllers

import (
	"errors"
	"fgd-golang/sesi-7/gorm/databases"
	"fgd-golang/sesi-7/gorm/models"
	"fmt"

	"gorm.io/gorm"
)

func GetUserById(id uint) {
	db := databases.GetDB()
	user := models.User{}

	err := db.First(&user, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("user not found")
			return
		}
		print("Error Finding User", err)
	}

	fmt.Printf("User Data : %+v\n", user)
}
