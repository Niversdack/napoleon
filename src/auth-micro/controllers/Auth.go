package controllers

import (
	"github.com/gin-gonic/gin"
	"napoleon/src/auth-micro/database"
	"napoleon/src/auth-micro/models"
	"napoleon/src/auth-micro/utils"
	"net/http"
)

// Auth godoc
// @Summary Auth
// @Description Auth Username, Password
// @Accept  json
// @Produce  json
// @Param user body RegRequest true "201 token"
// @Security ApiKeyAuth
// @Success 201
// @Router /auth [post]
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
