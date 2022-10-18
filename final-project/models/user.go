package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:"not null;unique;type:varchar(191)"`
	Username  string `gorm:"not null;unique; varchar(100);"`
	Password  string `gorm:"not null; varchar(255)"`
	Age       int    `gorm:"not null;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
