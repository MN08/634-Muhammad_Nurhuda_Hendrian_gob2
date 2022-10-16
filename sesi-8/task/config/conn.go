package config

import (
	"fgd-golang/sesi-8/task/models"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func DBInit() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/godb?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected")

	db.AutoMigrate(models.Person{})
	return db
}
