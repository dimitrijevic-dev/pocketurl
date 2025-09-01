package main

import (
	"pocketurl/persistence"
	"pocketurl/router"
)

func main() {
	persistence.Start()
	router.Start()
}
