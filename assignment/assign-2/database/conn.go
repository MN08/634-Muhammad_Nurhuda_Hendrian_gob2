package database

import (
	"fgd-golang/assignment/assign-2/models"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func StartDb() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/assign-order?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected")

	db.AutoMigrate(models.Item{}, models.Order{})
	return db
}
