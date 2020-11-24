package controllers

import (
	"github.com/gin-gonic/gin"
	"napoleon/src/todos-micro/database"
	"napoleon/src/todos-micro/models"
	"napoleon/src/todos-micro/utils"
	"net/http"
	"strings"
)

type DelRequest struct {
	Id uint `json:"id" binding:"required"`
}

func Delete(c *gin.Context) {
	var data DelRequest
	if err := c.Bind(&data); err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}
	token := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	claims := utils.GetClaims("napoleonJwT24", token)
	var todo models.Todo
	if database.DB.First(&todo).RecordNotFound() {
		c.Status(401)
		return
	}
	if todo.UserID != claims.ID {
		c.Status(401)
		return
	}
	database.DB.Unscoped().Delete(&todo)
	c.Status(201)
	return
}
