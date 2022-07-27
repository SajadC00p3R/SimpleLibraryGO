package main

import (
	"Library/Repository"
	"Library/Services"
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
	myRouter.MaxMultipartMemory = 8 << 20 // 8 MiB
	myRouter.GET("/allBooks", Services.GetBooks)
	myRouter.GET("/books/:id", Services.GetBook)
	myRouter.POST("/books/add", Services.NewBook)
	myRouter.PUT("/books/update/:id", Services.UpdateBook)
	myRouter.DELETE("/books/delete/:id", Services.DeleteBook)
	myRouter.POST("/order/neworder", Services.NewOrder)
	myRouter.PUT("/order/:order_id", Services.UpdateOrderStatus)
	myRouter.DELETE("/order/:order_id", Services.DeleteOrder)
	myRouter.GET("/order/:user_id", Services.GetOrderByUser)
	myRouter.GET("/order/update/:Status", Services.GetOrderByStatus)
	myRouter.GET("/order/book_id", Services.GetOrderByBook)
	myRouter.POST("/user/uploadAvatar", Services.UploadAvatar)
	myRouter.POST("/uploadImage", Services.UploadAvatar)
	myRouter.POST("/user/addnew", Services.SignUp)
	myRouter.POST("/user/update/:user_id", Services.UpdateUser)
	log.Fatal(http.ListenAndServe(":1000", myRouter))
}
