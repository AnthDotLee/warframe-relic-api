package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()

	relicsApi := router.Group("relics")
	{
		relicsApi.GET("/", func(c *gin.Context) {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusOK, Relics)
		})
	}
}
