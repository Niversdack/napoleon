package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Todo struct {
	gorm.Model
	UserID      uint
	Name        string
	Description string
	EndTime     time.Time
}
