package Entity

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	OrderId string `json:"OrderId" `
	UserId  string `json:"UserId" gorm:"foreignKey"`
	BookId  string `json:"BookId" gorm:"foreignKey"`
	Status  string `json:"Status" gorm:"foreignKey"`
}
