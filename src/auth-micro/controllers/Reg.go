package controllers

import (
	"github.com/gin-gonic/gin"
	"napoleon/src/auth-micro/database"
	"napoleon/src/auth-micro/models"
	"net/http"
)

type RegRequest struct {
	Username string `json:"username" binding:"required" example:"PetrovVasya"`
	Password string `json:"password" binding:"required" example:"DomPushkina"`
}

// Reg godoc
// @Summary Reg
// @Description Reg by Username, Password
// @Accept  json
// @Produce  json
// @Param user body RegRequest true "code 201"
// @Security ApiKeyAuth
// @Success 201
// @Router /reg [post]
func Reg(c *gin.Context) {
	var data RegRequest

	if err := c.Bind(&data); err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}
	User := models.User{
		Username: data.Username,
		Password: data.Password,
	}
	database.DB.Create(&User)
	c.Status(201)
	return
}
