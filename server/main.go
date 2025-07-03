package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	rdb := redis.NewClient(&redis.Options{Addr: "redis_db:6379"})

	dsn := "host=go_db user=postgres password=password dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect db: %v \n", err)
		return
	}

	err = db.Exec("CREATE TABLE IF NOT EXISTS NUMBERS(value int)").Error
	if err != nil {
		fmt.Printf("failed to create values table: %v\n", err)
		return
	}

	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/values/all", func(c *gin.Context) {
		// postgres
		var values []struct {
			Value int `gorm:"column:value"`
		}
		err := db.Raw("SELECT * FROM NUMBERS").Scan(&values).Error
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		var vals []int = []int{}
		for _, v := range values {
			vals = append(vals, v.Value)
		}

		c.JSON(200, gin.H{
			"values": vals,
		})
	})

	router.GET("/values/current", func(c *gin.Context) {
		// redis
		values, err := rdb.HGetAll(c, "numbers").Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"values": values,
		})
	})

	router.POST("/values", func(c *gin.Context) {
		var req struct {
			Value int `json:"value"`
		}
		c.ShouldBind(&req)

		err := db.Exec(`INSERT INTO NUMBERS(value) VALUES (?)`, req.Value).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		_, err = rdb.HSet(c, "numbers", map[string]interface{}{fmt.Sprintf("%v", req.Value): req.Value}).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{"working": true, "value": req})
	})

	router.Run()
}
