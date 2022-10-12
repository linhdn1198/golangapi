package main

import (
	"golangapi/routes"
)

func main() {
	server := routes.CreateRouter()
	server.Run()
}