package database

import (
	"fmt"

	"github.com/backsoul/doberman/pkg/services"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (db *gorm.DB) {
	dsn := services.Get("DB")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		services.Log(fmt.Sprint(err), "Error")
	}
	return db
}
