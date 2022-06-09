package main

import (
	"testing"

	"github.com/backsoul/doberman/internal/database"
	"github.com/backsoul/doberman/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {

	t.Run("Test Create Product - CreateProduct", func(t *testing.T) {
		db := database.InitDB()
		var articles = []types.Article{{Title: "title", Content: "content example jeje"}}
		assert.Equal(t, db.Create(&articles).Error, nil)
	})

}
