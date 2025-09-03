package router

import (
	"math/rand"
	"pocketurl/persistence"
	"time"
)

func GenerateLink(linkRequest linkRequest) persistence.Link {
	newLink := persistence.Link{
		OriginUrl:      generateRoute(),
		ExpiresAt:      generateExpirationDate(),
		DestinationUrl: linkRequest.DestinationUrl,
		Domain:         linkRequest.Domain,
	}
	return newLink
}

func generateExpirationDate() time.Time {
	return time.Now().AddDate(0, 0, 7)
}

func generateRoute() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := ""
	for range 6 {
		b += string(charset[rand.Intn(len(charset))])
	}
	return b
}
