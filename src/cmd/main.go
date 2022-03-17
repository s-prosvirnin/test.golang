package main

import "github.com/gin-gonic/gin"
import "log"

func main() {
	log.Println("------19------")
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "halo",
		})
	})
	r.Run(":3001")
}
