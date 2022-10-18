package models

import "time"

type Photo struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"not null;type:varchar(191)"`
	Caption   string `gorm:"not null; varchar(255);"`
	PhotoUrl  string `gorm:"not null; varchar(255)"`
	UserID    uint   `json:"UserID" gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
