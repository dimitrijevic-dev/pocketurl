package router

import (
	"fmt"
	"math/rand"
)

func GenerateLink(linkRequest linkRequest) Link {
	randomRoute := generateRoute()
	fmt.Println(randomRoute)
	newLink := Link{
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