package router

import (
	"net/http"
	"pocketurl/config"
	"pocketurl/persistence"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type linkRequest struct {
	DestinationUrl string `json:"destination_url"`
	Domain         string `json:"domain"`
}

func Start() {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	router.GET("/:origin", getLinkByOrigin)
	router.POST("/links", postLinks)

	certFile := "/etc/letsencrypt/live/api.pocketurl.zip/fullchain.pem"
	keyFile := "/etc/letsencrypt/live/api.pocketurl.zip/privkey.pem"
	_, _ = certFile, keyFile

	port := config.GetEnv("PORT")
	if port == "" {
		port = "443"
	}

	//router.Run("localhost:8080")
	router.RunTLS(":"+port, certFile, keyFile)
}

func postLinks(c *gin.Context) {
	var newLinkRequest linkRequest
	if err := c.BindJSON(&newLinkRequest); err != nil {
		return
	}
	if !(strings.HasPrefix(newLinkRequest.DestinationUrl, "https://") || strings.HasPrefix(newLinkRequest.DestinationUrl, "http://")) {
		newLinkRequest.DestinationUrl = "https://" + newLinkRequest.DestinationUrl
	}

	newLink := GenerateLink(newLinkRequest)
	for {
		existingLink := persistence.GetLinkByOrigin(newLink.OriginUrl)
		if existingLink == nil || existingLink.ID == 0 {
			break
		}
		newLink = GenerateLink(newLinkRequest)
	}

	if err := persistence.AddLink(newLink); err != nil {
		return
	}
	c.IndentedJSON(http.StatusCreated, newLink)
}

func getLinkByOrigin(c *gin.Context) {
	originParameter := c.Param("origin")
	result := persistence.GetLinkByOrigin(originParameter)
	if result == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "link not found"})
	} else if result.ExpiresAt.After(time.Now()) {
		c.Redirect(http.StatusPermanentRedirect, result.DestinationUrl)
	} else {
		persistence.DeleteLink(*result)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "link not found"})
	}
}
