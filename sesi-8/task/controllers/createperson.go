package controllers

import (
	"fgd-golang/sesi-8/task/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (idb *InDB) CreatePerson(c *gin.Context) {
	var (
		person models.Person
		result gin.H
	)

	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
	person.First_Name = first_name
	person.Last_Name = last_name
	idb.DB.Create(&person)

	result = gin.H{
		"result": person,
	}
	c.JSON(http.StatusOK, result)

}
