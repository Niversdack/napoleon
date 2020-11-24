package controllers

import (
	"github.com/gin-gonic/gin"
	"napoleon/src/auth-micro/database"
	"napoleon/src/auth-micro/models"
	"napoleon/src/auth-micro/utils"
	"net/http"
)

func Auth(c *gin.Context) {
	var data RegRequest

	if err := c.Bind(&data); err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}
	var user models.User
	if database.DB.Where("username = ?", data.Username).First(&user).RecordNotFound() {
		c.Status(http.StatusUnauthorized)
		return
	}

	accesstoken := utils.GenerateJWT("napoleonJwT24", user.Username, user.ID)

	c.JSON(201, gin.H{"token": accesstoken})
	return
}
