package models

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ID          uint   `gorm:"primaryKey"`
	Code        string `gorm:"not null;type:varchar(10)"`
	Description string `gorm:"not null;type:varchar(191)"`
	Quantity    uint
	OrderID     uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (i *Item) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Item Before Created()")

	if len(i.Code) < 4 {
		err = errors.New("item code is too short")
	}
	return
}
