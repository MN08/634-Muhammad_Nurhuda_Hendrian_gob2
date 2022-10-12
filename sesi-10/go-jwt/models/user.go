package models

import (
	"fgd-golang/sesi-10/go-jwt/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModels
	FullName string    `gorm:"not null" json:"full_name" form:"full_name" valid:"required"`
	Email    string    `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required"`
	Password string    `gorm:"not null" json:"password" form:"password" valid:"required"`
	Products []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`
}

func (u *User) beforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPassword(u.Password)
	err = nil
	return

}
