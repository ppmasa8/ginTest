package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/darksoul", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H {
			"message": "You Died",
		})
	})
	r.Run()
}