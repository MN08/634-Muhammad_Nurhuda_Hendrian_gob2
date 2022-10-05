package main

import (
	userControllers "fgd-golang/sesi-7/gorm/controllers/users"
	"fgd-golang/sesi-7/gorm/databases"
)

func main() {
	databases.StartDb()
	//create User to Db
	userControllers.CreateUser("andre@mjd.online")
	//get single user from database
	userControllers.GetUserById(1)
	//update User data
	userControllers.UpdateUser(1, "andrea@mjd.online")
}
