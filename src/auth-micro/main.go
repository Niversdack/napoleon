package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"napoleon/src/auth-micro/controllers"
	"napoleon/src/auth-micro/database"
	"napoleon/src/auth-micro/models"
	"napoleon/src/auth-micro/utils"
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
	r.POST("/reg", controllers.Reg)
	r.POST("/auth", controllers.Auth)
	r.POST("/user", controllers.Get).Use(AuthMiddleware())
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
