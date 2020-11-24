package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"napoleon/src/auth-micro/controllers"
	"napoleon/src/auth-micro/database"
	"napoleon/src/auth-micro/docs"
	"napoleon/src/auth-micro/models"
	"napoleon/src/auth-micro/utils"
)

//@securityDefinitions.apikey ApiKeyAuth
//@in header
//@name Authorization
func main() {

	docs.SwaggerInfo.Title = "Napoleon Test API"

	docs.SwaggerInfo.Schemes = []string{"http"}
	db := database.Initialize()
	database.Migrate()
	defer db.Close()

	r := gin.Default()
	r.POST("/reg", controllers.Reg)
	r.POST("/auth", controllers.Auth)

	auth := r.Group("/")
	auth.Use(AuthMiddleware())
	{
		auth.POST("/user", controllers.Get)
	}
	r.GET("/swagger/auth-micro/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header["Authorization"][0]
		claims := utils.GetClaims("napoleonJwT24", token)

		var user models.User
		database.DB.Where("username=?", claims.Username).First(&user)
		if user.ID != claims.ID {
			c.Status(401)
		}
		c.Next()
	}
}
