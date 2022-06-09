package types

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title   string
	Content string
}
