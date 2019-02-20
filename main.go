package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()

	relicsApi := router.Group("/relics")
	{
		relicsApi.GET("/", func(c *gin.Context) {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusOK, getRelics())
		})
	}
	relicSearchApi := router.Group("/relic")
	{
		relicSearchApi.GET("/:era/:name", relicSearchHandler)
	}

	err := router.Run(":3000")

	if err != nil {
		fmt.Printf("Server was unable to start do to error: %s", err)
	}
}

func relicSearchHandler(c *gin.Context) {

}
