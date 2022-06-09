package main

import (
	"fmt"

	"github.com/backsoul/doberman/pkg/services"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title   string
	Content string
}

func main() {

	dsn := services.Get("DB")
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		services.Log("failed to connect database", "Error")
	}
	db.AutoMigrate(&Article{})
}
