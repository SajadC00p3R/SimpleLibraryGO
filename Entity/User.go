package Entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name   string `json:"userName" `
	UserId string `json:"userId" `
	Email  string `json:"userEmail" `
	Age    int    `json:"Age" `
	Avatar string `json:"Avatar" `
}
