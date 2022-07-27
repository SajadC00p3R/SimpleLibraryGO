package Services

import (
	"Library/Entity"
	"Library/Repository"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type newUser struct {
	Name   string `json:"userName" binding:"required"`
	UserId string `json:"userId" binding:"required"`
	Email  string `json:"userEmail" binding:"required"`
	Age    int    `json:"Age" binding:"required"`
	Avatar string `json:"Avatar" binding:"required"`
}
type updateUser struct {
	Name   string `json:"userName" `
	UserId string `json:"userId"`
	Email  string `json:"userEmail"`
	Age    int    `json:"Age" `
	Avatar string `json:"Avatar"`
}

func SignUp(c *gin.Context) {
	var user newUser

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newuser := Entity.User{Name: user.Name, UserId: user.UserId, Email: user.Email, Age: user.Age, Avatar: user.Avatar}

	db, err := Repository.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Create(&newuser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

/*
func getUserData()  {

}
//todo: hendle tokens before handle this function
*/

func UpdateUser(c *gin.Context) {

	var user Entity.User

	db, err := Repository.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Where("user_id = ?", c.Param("user_id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found!"})
		return
	}

	var userUpdate updateUser

	if err := c.ShouldBindJSON(&userUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&user).Updates(Entity.User{Name: userUpdate.Name, UserId: userUpdate.UserId, Email: userUpdate.Email, Age: userUpdate.Age, Avatar: userUpdate.Avatar}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
