package ServiceBook

import (
	"Library/Entity"
	"Library/Repository"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type newBook struct {
	BookId string `json:"Id",gorm:"primaryKey" binding:"required"`
	Title  string `json:"Title" json:"title" binding:"required"`
	Pub    string `json:"pub" json:"pub" binding:"required"`
	Year   string `json:"year" json:"year" binding:"required"`
	Author string `json:"author" json:"author" binding:"required"`
}

type BookUpdate struct {
	BookId string `json:"Id"`
	Title  string `json:"Title"`
	Pub    string `json:"pub"`
	Year   string `json:"year"`
	Author string `json:"author"`
}

func GetBooks(c *gin.Context) {

	var groceries []Entity.Book

	db, err := Repository.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Find(&groceries).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, groceries)
}
func GetBook(c *gin.Context) {

	var grocery Entity.Book

	db, err := Repository.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Where("id= ?", c.Param("id")).First(&grocery).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, grocery)

}

func NewBook(c *gin.Context) {

	var book newBook

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newBook := Entity.Book{BookId: book.BookId, Title: book.Title, Pub: book.Pub, Year: book.Year, Author: book.Author}

	db, err := Repository.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Create(&newBook).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newBook)
}
func UpdateBook(c *gin.Context) {

	var book Entity.Book

	db, err := Repository.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found!"})
		return
	}

	var bookUpdate BookUpdate

	if err := c.ShouldBindJSON(&bookUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&book).Updates(Entity.Book{
		BookId: bookUpdate.BookId, Title: bookUpdate.Title, Pub: bookUpdate.Pub, Year: bookUpdate.Year, Author: bookUpdate.Author}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)

}
func DeleteBook(c *gin.Context) {

	var book Entity.Book

	db, err := Repository.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found!"})
		return
	}

	if err := db.Delete(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})

}
