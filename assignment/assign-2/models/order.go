package models

import (
	"time"
)

type Order struct {
	ID        uint   `gorm:"primaryKey"`
	Cust_Name string `gorm:"not null;type:varchar(191)"`
	OrderedAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
