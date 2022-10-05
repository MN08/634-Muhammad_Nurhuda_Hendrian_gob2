package userControllers

import (
	"fgd-golang/sesi-7/gorm/databases"
	"fgd-golang/sesi-7/gorm/models"
	"fmt"
)

func UpdateUser(id uint, email string) {
	db := databases.GetDB()
	user := models.User{}
	err := db.Model(&user).Where("id = ?", id).Updates(models.User{Email: email}).Error

	if err != nil {
		fmt.Println("Error updating user data:", err)
		return
	}

	fmt.Printf("User Email Updated: %+v\n", user.Email)

}
