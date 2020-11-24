package controllers

import (
	"github.com/gin-gonic/gin"
	"napoleon/src/todos-micro/database"
	"napoleon/src/todos-micro/models"
	"napoleon/src/todos-micro/utils"
	"net/http"
)

type GetByIDRequest struct {
	ID uint `json:"id"`
}

// GetByID godoc
// @Summary GetByID
// @Description Get by ID
// @Accept  json
// @Produce  json
// @Param todos body GetByIDRequest true "{object} models.Todo"
// @Security ApiKeyAuth
// @Success 201 {object} models.Todo
// @Router /getbyid [post]
func GetByID(c *gin.Context) {
	var data GetByIDRequest
	if err := c.Bind(&data); err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}
	token := c.Request.Header["Authorization"][0]
	claims := utils.GetClaims("napoleonJwT24", token)
	var todo models.Todo
	if database.DB.First(&todo).RecordNotFound() {
		c.JSON(201, gin.H{})
		return
	}
	if todo.UserID != claims.ID {
		c.Status(401)
		return
	}
	c.JSON(201, todo)
	return
}

type GetByTimeRequest struct {
	Time int64 `json:"time"`
}

// GetByTime godoc
// @Summary GetByTime
// @Description Get by Time
// @Accept  json
// @Produce  json
// @Param todos body GetByTimeRequest true "{array} models.Todo"
// @Security ApiKeyAuth
// @Success 201 {array} models.Todo
// @Router /getbytime [post]
func GetByTime(c *gin.Context) {
	var data GetByTimeRequest
	if err := c.Bind(&data); err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}
	token := c.Request.Header["Authorization"][0]
	claims := utils.GetClaims("napoleonJwT24", token)
	var todos []models.Todo

	if database.DB.Where("user_id=? AND end_time<=(SELECT TIMESTAMP 'epoch' + ? * INTERVAL '1 second')", claims.ID, data.Time).Order("end_time asc").Find(&todos).RecordNotFound() {
		c.JSON(201, gin.H{})
		return
	}
	c.JSON(201, todos)
	return
}

// GetAll godoc
// @Summary GetAll
// @Description Get by Time
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 201 {array} models.Todo
// @Router /getall [post]
func GetAll(c *gin.Context) {
	token := c.Request.Header["Authorization"][0]
	claims := utils.GetClaims("napoleonJwT24", token)
	var todos []models.Todo

	if database.DB.Where("user_id=?", claims.ID).Order("end_time asc").Find(&todos).RecordNotFound() {
		c.JSON(201, gin.H{})
		return
	}
	c.JSON(201, todos)
	return
}
