package repository

import (
	"fgd-golang/final-project/models"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func StartDb() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/finalproject?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected")

	db.AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})
	return db
}
