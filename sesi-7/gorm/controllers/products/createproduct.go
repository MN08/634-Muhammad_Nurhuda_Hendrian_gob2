package productControllers

import (
	"fgd-golang/sesi-7/gorm/databases"
	"fgd-golang/sesi-7/gorm/models"
	"fmt"
)

func CreateProduct(userId uint, brand string, name string) {
	db := databases.GetDB()

	Product := models.Product{
		UserID: userId,
		Brand:  brand,
		Name:   name,
	}

	err := db.Create(&Product).Error
	if err != nil {
		fmt.Println("Error creating product data:", err.Error())
		return
	}

	fmt.Println("new product created successfully :", Product)
}
