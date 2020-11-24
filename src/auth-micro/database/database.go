package database

import (
	"time"

	"dreamteam-api/src/auth-micro/common"
	"dreamteam-api/src/auth-micro/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

//var connectionSting string

func Initialize() *gorm.DB {

	db, err := gorm.Open("postgres", "host="+common.GlobalConfig.PostgresDb.Host+" port="+common.GlobalConfig.PostgresDb.Port+" user="+common.GlobalConfig.PostgresDb.Username+" password="+common.GlobalConfig.PostgresDb.Password+" dbname="+common.GlobalConfig.PostgresDb.Name+" sslmode="+common.GlobalConfig.PostgresDb.Sslmode)

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

	DB.AutoMigrate(&models.Token{}, &models.User{}, &models.Picture{}, &models.Params{}, &models.Online{})

	log.Info("Table was successfully created")

}
