package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type link struct {
	OriginUrl string `json:"origin"`
	Duration int `json:"duration"`
	DestinationUrl string `json:"destination"`
}

var links = []link{
		{OriginUrl: "test", Duration: 3600, DestinationUrl: "wikipedia.com"},
	}

func Start() {
	router := gin.Default()
	router.GET("/links", getLinks)
	router.GET("/links/:origin", getLinkByOrigin)
	router.POST("/links", postLinks)

	router.Run("localhost:8080")
}

func getLinks(c *gin.Context) {
	fmt.Println("Returning all...")
	c.IndentedJSON(http.StatusOK, links)
}

func postLinks(c *gin.Context) {
	var newLink link

	if err := c.BindJSON(&newLink); err != nil { return }

	links =  append(links, newLink)
	c.IndentedJSON(http.StatusCreated, newLink)
}

func getLinkByOrigin(c *gin.Context) {
	originParameter := c.Param("origin")
	fmt.Println(originParameter)

	for _, a := range links {
		if a.OriginUrl == originParameter {
			c.IndentedJSON(http.StatusOK, a) 
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message":"link not found"})
}