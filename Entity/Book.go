package Entity

import "gorm.io/gorm"

type Book struct {
	gorm.Model

	BookId string `json:"Id" gorm:"primaryKey"`
	Title  string `json:"Title" json:"title"`
	Pub    string `json:"pub" json:"pub"`
	Year   string `json:"year" json:"year"`
	Author string `json:"author" json:"author"`
}
