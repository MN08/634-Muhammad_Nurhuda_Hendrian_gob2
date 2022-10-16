package models

type Item struct {
	ID          int    `json:"itemId" gorm:"primaryKey;autoIncrement:true"`
	Code        string `json:"itemCode" gorm:"not null;type:varchar(191)"`
	Description string `json:"description" gorm:"not null;type:varchar(191)"`
	Quantity    int    `json:"quantity"`
	OrderID     int    `json:"OrderID" gorm:"index"`
	Items       []Item `json:"item" gorm:"foreignKey:OrderID"`
}
