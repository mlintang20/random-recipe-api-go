package main

import (
	"net/http"
	"webapp1/handler"

	"github.com/gin-gonic/gin"
)

// buat service yg mengembalikan resep (nama dan bahan2) dan harga total yg dibutuhkan:
// 1. ambil data dari https://api.spoonacular.com/recipes/random
// 2. bikin server
// 3. bikin endpoint dan handler nya

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"data": nil,
		})
	})

	r.GET("/recipes", handler.HandlerGetRecipe)

	r.Run(":8080")
}