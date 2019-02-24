package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
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
	searchEra := strings.ToLower(c.Param("era"))
	searchName := strings.ToLower(c.Param("name"))

	for i := 0; i < len(Relics); i++ {
		if strings.ToLower(Relics[i].Era) == searchEra && strings.ToLower(Relics[i].Name) == searchName {
			c.JSON(http.StatusOK, Relics[i])
			return
		}
	}
	c.AbortWithStatus(http.StatusNotFound)
}
