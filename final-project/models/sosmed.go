package models

import "time"

type SocialMedia struct {
	ID               uint   `gorm:"primaryKey"`
	Name             string `gorm:"not null; varchar(255);"`
	social_media_url string `gorm:"not null; varchar(255);"`
	UserID           uint   `json:"UserID" gorm:"index"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
