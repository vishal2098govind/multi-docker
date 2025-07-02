package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/values/all", func(c *gin.Context) {
		// postgres
		c.JSON(200, gin.H{
			"values": []int{1, 2, 3, 4},
		})
	})

	router.GET("/values/current", func(c *gin.Context) {
		// redis
		c.JSON(200, gin.H{
			"values": map[int]int{
				1: 1,
				2: 2,
				3: 3,
			},
		})
	})

	router.POST("/values", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"working": true,
		})
	})

	router.Run()
}
