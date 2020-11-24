package database

import (
	//"napoleon/src/auth-micro/configs"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	"napoleon/src/auth-micro/models"
)

var DB *gorm.DB

//var connectionSting string

func Initialize() *gorm.DB {

	db, err := gorm.Open("postgres", "host=db dbname=postgres sslmode=disable password=H6bLerShrQ user=postgres")

	if err != nil {
		log.Fatal(err)
	}

	db.DB().SetMaxIdleConns(10)             // The default is defaultMaxIdleConns (= 2)
	db.DB().SetMaxOpenConns(1000)           // The default is 0 (unlimited)
	db.DB().SetConnMaxLifetime(time.Second) // The default is 0 (connections reused forever)

	//db.LogMode(true)
	DB = db

	log.Info("Database connected")
	return DB
}

// Migrate all needed tables
func Migrate() {

	DB.AutoMigrate(&models.User{})

	log.Info("Table was successfully created")

}
