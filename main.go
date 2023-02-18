package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Recipe struct {
	Name         string    `json:"name"`
	Tags         []string  `json:"tags"`
	Ingredients  []string  `json:"igredients"`
	Instructions []string  `json:"instructions"`
	PublishedAt  time.Time `json:"publishedAt"`
}

func IndexHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome",
	})
}

func main() {
	router := gin.Default()
	router.GET("/", IndexHandler)
	router.Run(":8080")
}
