package controllers

import (
	"github.com/gin-gonic/gin"
	"napoleon/src/auth-micro/database"
	"napoleon/src/auth-micro/models"
	"net/http"
)

type RequestUser struct {
	Id uint `json:"id"`
}

func Get(c *gin.Context) {
	var data RequestUser
	if err := c.Bind(&data); err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}
	var user models.User
	database.DB.First(&user, data.Id)
	c.JSON(201, user.ToWithoutPassword())
	return
}
