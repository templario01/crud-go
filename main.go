package main

import (
	router "crud-go/src/infrastructure"
)

func main() {
	router := router.SetupRouter()
	router.Run(":8080")
}
