package main

import (
	"Library/Repository"
	"Library/ServiceBook"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	handleDB()
	handleRequests()

}

func handleDB() {
	db, err := Repository.Database()
	if err != nil {
		log.Println(err)
	}
	_, _ = db.DB()
}
func handleRequests() {
	myRouter := gin.Default()
	myRouter.GET("/allBooks", ServiceBook.GetBooks)
	myRouter.GET("/books/:id", ServiceBook.GetBook)
	myRouter.POST("/books/add", ServiceBook.NewBook)
	myRouter.PUT("/books/update/:id", ServiceBook.UpdateBook)
	myRouter.DELETE("/books/delete/:id", ServiceBook.DeleteBook)
	log.Fatal(http.ListenAndServe(":1000", myRouter))
}
