package models

import (
	"time"
)

type Order struct {
	ID           int       `json:"orderId" gorm:"primaryKey;autoIncrement:true"`
	CustomerName string    `json:"customerName" gorm:"not null;type:varchar(191)"`
	OrderedAt    time.Time `json:"orderedAt"`
}
