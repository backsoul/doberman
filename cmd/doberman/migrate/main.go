package main

import (
	"github.com/backsoul/doberman/internal/database"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title   string
	Content string
}

func main() {
	db := database.InitDB()
	db.AutoMigrate(&Article{})
}
