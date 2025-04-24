package main

import "albumapi/api"

func main() {
	server := api.NewAPIServer(":8090")
	server.Run()
}
