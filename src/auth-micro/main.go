package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"napoleon/src/auth-micro/configs"
	"napoleon/src/auth-micro/controllers"
	"napoleon/src/auth-micro/database"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}
func main() {
	flag.Parse()
	config := configs.NewConfig()

	db := database.Initialize()
	database.Migrate()
	defer db.Close()

	r := gin.Default()
	r.POST("/reg", controllers.Reg)
	r.POST("/auth", controllers.Auth)
	err := r.Run(config.BindAddr)
	if err != nil {
		log.Fatal(err)
	}
}
