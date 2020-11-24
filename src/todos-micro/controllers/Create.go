package controllers

import (
	"github.com/gin-gonic/gin"
	"napoleon/src/todos-micro/database"
	"napoleon/src/todos-micro/models"
	"napoleon/src/todos-micro/utils"
	"net/http"
	"strings"
	"time"
)

type CreateRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	EndTime     int64  `json:"end_time" binding:"required"`
}

func Create(c *gin.Context) {
	var data CreateRequest
	if err := c.Bind(&data); err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}
	token := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	claims := utils.GetClaims("napoleonJwT24", token)
	todo := models.Todo{
		UserID:      claims.ID,
		Name:        data.Name,
		Description: data.Description,
		EndTime:     time.Unix(data.EndTime, 0),
	}
	database.DB.Create(&todo)
	c.Status(201)
	return
}
