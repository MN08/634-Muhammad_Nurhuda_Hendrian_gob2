package models

import "time"

type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `json:"UserID" gorm:"index"`
	PhotoID   uint   `json:"PhotoID" gorm:"index"`
	Comment   string `gorm:"varchar(255);"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
