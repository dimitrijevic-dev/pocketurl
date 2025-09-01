package router

import (
	"math/rand"
	"pocketurl/persistence"
)

func GenerateLink(linkRequest linkRequest) persistence.Link {
	randomRoute := generateRoute()
	newLink := persistence.Link{
		OriginUrl: randomRoute,
		ExpiresAt: linkRequest.ExpiresAt,
		DestinationUrl: linkRequest.DestinationUrl,
	}
	return newLink
}

func generateRoute() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := ""
	for range 6 {
		b += string(charset[rand.Intn(len(charset))])
	}
	return b
}