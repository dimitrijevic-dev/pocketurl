package main

import (
	"fmt"
	"pocketurl/persistence"
	"pocketurl/router"
)

func main() {
	persistence.Connect()
	fmt.Println("Hello - " + persistence.GetDestinationByOrigin("A9mx3F"))
	router.Start()
}
