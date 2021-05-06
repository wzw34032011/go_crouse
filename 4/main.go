package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	g := gin.Default()
	g.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	err := g.Run("127.0.0.1:8080")
	if err != nil {
		log.Fatalf("应用启动异常，%s", err.Error())
	}
}
