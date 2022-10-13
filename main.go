package main

import (
	"golangapi/database/seeders"
	"golangapi/routes"
	"golangapi/util"
	"os"
)

func main() {
	util.LoadEnv()

	if len(os.Args) > 1 && os.Args[1] == "Seeder" {
		seeders.Run()
	} else {
		server := routes.CreateRouter()
		server.Run(":" + os.Getenv("APP_PORT"))
	}
}