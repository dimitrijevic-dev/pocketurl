package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type link struct {
	originUrl string `json"origin"`
	duration int `json"duration"`
	destinationUrl string `json"destination"`
}

var links = []link{
		{originUrl: "test", duration: 3600, destinationUrl: "wikipedia.com"},
	}

func main() {
	router := gin.Default()
	router.GET("/links", getLinks)

	router.Run("localhost:8080")
}

func getLinks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, links)
}