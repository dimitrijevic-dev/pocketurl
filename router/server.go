package router

import (
	"net/http"
	"pocketurl/persistence"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type linkRequest struct {
	ExpiresAt time.Time `json:"expires_at"`
	DestinationUrl string `json:"destination_url"`
}

func Start() {
	router := gin.Default()
	router.GET("/:origin", getLinkByOrigin)
	router.POST("/links", postLinks)

	router.Run("localhost:8080")
}

func postLinks(c *gin.Context) {
	var newLinkRequest linkRequest
	if err := c.BindJSON(&newLinkRequest); err != nil { return }
	if !(strings.HasPrefix(newLinkRequest.DestinationUrl,"https://") || strings.HasPrefix(newLinkRequest.DestinationUrl,"http://")) {
		newLinkRequest.DestinationUrl = "https://"+newLinkRequest.DestinationUrl
	}

	newLink := GenerateLink(newLinkRequest)
	if err := persistence.AddLink(newLink); err != nil { return }
	c.IndentedJSON(http.StatusCreated, newLink)
}

func getLinkByOrigin(c *gin.Context) {
	originParameter := c.Param("origin")
	result := persistence.GetLinkByOrigin(originParameter)
	if result == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message":"link not found"})
	} else if result.ExpiresAt.After(time.Now()){
		c.Redirect(http.StatusPermanentRedirect,result.DestinationUrl)
	} else {
		persistence.DeleteLink(*result)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message":"link not found"})
	}
}