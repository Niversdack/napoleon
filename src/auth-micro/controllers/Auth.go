package controllers

import (
	"github.com/gin-gonic/gin"
	"napoleon/src/auth-micro/configs"
	"napoleon/src/auth-micro/database"
	"napoleon/src/auth-micro/models"
	"napoleon/src/auth-micro/utils"
	"net/http"
	"time"
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
	accesstoken := utils.GenerateJWT(configs.GlobalConfig.SecretKey, 99999*time.Hour, data.Username)
	c.JSON(201, gin.H{"token": accesstoken})
	return
}
