package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"napoleon/src/todos-micro/controllers"
	"napoleon/src/todos-micro/database"
	"napoleon/src/todos-micro/docs"
	"napoleon/src/todos-micro/models"
	"napoleon/src/todos-micro/utils"
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
	r.Use(gin.Recovery())
	auth := r.Group("/")
	auth.Use(AuthMiddleware())
	{
		auth.POST("/create", controllers.Create)
		auth.POST("/delete", controllers.Delete)
		auth.POST("/update", controllers.Update)
		auth.POST("/getall", controllers.GetAll)
		auth.POST("/getbyid", controllers.GetByID)
		auth.POST("/getbytime", controllers.GetByTime)
	}
	r.GET("/swagger/todos-micro/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

//test gitflow
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header["Authorization"][0]
		claims := utils.GetClaims("napoleonJwT24", token)
		var asdasdas int
		_ = asdasdas
		_ = asdasdas

		_ = asdasdas

		var user models.User
		database.DB.Where("username=?", claims.Username).First(&user)
		if user.ID != claims.ID {
			c.Status(401)
		}
		c.Next()

	}
}
