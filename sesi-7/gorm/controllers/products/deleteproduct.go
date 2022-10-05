package productControllers

import (
	"fgd-golang/sesi-7/gorm/databases"
	"fgd-golang/sesi-7/gorm/models"
	"fmt"
)

func DeleteProductById(id uint) {
	db := databases.GetDB()

	Product := models.Product{}

	err := db.Where("id = ?", id).Delete(&Product).Error
	if err != nil {
		fmt.Println("Error deletes product data:", err.Error())
		return
	}

	fmt.Printf("id %d product successfully deleted", id)
}
