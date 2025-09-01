package router

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Link struct {
	ID int8 `json:"id"`
	OriginUrl string `json:"origin_url"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
	DestinationUrl string `json:"destination_url"`
}

type linkRequest struct {
	ExpiresAt time.Time `json:"expires_at"`
	DestinationUrl string `json:"destination_url"`
}

var links = []Link{
		{OriginUrl: "test", ExpiresAt: time.Now(), DestinationUrl: "wikipedia.com"},
	}

func Start() {
	router := gin.Default()
	router.GET("/links", getLinks)
	router.GET("/:origin", getLinkByOrigin)
	router.POST("/links", postLinks)

	router.Run("localhost:8080")
}

func getLinks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, links)
}

func postLinks(c *gin.Context) {
	var newLinkRequest linkRequest
	if err := c.BindJSON(&newLinkRequest); err != nil { return }

	newLink := GenerateLink(newLinkRequest)
	links = append(links, newLink)
	c.IndentedJSON(http.StatusCreated, newLink)
}

func getLinkByOrigin(c *gin.Context) {
	originParameter := c.Param("origin")
	for _, a := range links {
		if a.OriginUrl == originParameter {
			c.IndentedJSON(http.StatusOK, a) 
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message":"link not found"})
}