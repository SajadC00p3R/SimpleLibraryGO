package Services

import (
	"Library/Entity"
	"Library/Repository"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type newOrder struct {
	OrderId string `json:"OrderId" gorm:"primaryKey" binding:"required"`
	UserId  string `json:"UserId" gorm:"foreignKey" binding:"required"`
	BookId  string `json:"BookId" gorm:"foreignKey" binding:"required"`
	Status  string `json:"Status" binding:"required"`
}

type UpdateOrder struct {
	Status string `json:"Status" binding:"required"`
}

func NewOrder(c *gin.Context) {
	var order newOrder
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newOrder := Entity.Order{OrderId: order.OrderId, UserId: order.UserId, BookId: order.BookId, Status: order.Status}

	db, err := Repository.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Create(&newOrder).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newOrder)
}

func UpdateOrderStatus(c *gin.Context) {
	var order Entity.Order
	db, err := Repository.Database()
	if err != nil {
		log.Println(err)
	}
	if err := db.Where("order_id = ?", c.Param("order_id")).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "order not found!"})
		return
	}
	var updateOrder UpdateOrder

	if err := c.ShouldBindJSON(&updateOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Model(&order).Updates(Entity.Order{OrderId: order.OrderId, UserId: order.UserId, BookId: order.BookId, Status: updateOrder.Status}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, order)
}

func deleteOrder()      {}
func getOrderByUser()   {}
func getOrderByStatus() {}
func getOrderByBook()   {}
