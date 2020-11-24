package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"napoleon/src/todos-micro/controllers"
	"napoleon/src/todos-micro/database"
	"napoleon/src/todos-micro/models"
	"napoleon/src/todos-micro/utils"
	"strings"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}
func main() {
	flag.Parse()

	db := database.Initialize()
	database.Migrate()
	defer db.Close()

	r := gin.Default()
	r.POST("/create", controllers.Create)
	r.POST("/delete", controllers.Delete)
	r.POST("/update", controllers.Update)
	r.POST("/getall", controllers.GetAll)
	r.POST("/getbyid", controllers.GetByID)
	r.POST("/getbytime", controllers.GetByTime)

	r.Use(AuthMiddleware())

	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
		claims := utils.GetClaims("napoleonJwT24", token)

		var user models.User
		database.DB.Where("username=?", claims.Username).First(&user)
		if user.ID != claims.ID {
			c.Status(401)
		}
		c.Next()
	}
}
