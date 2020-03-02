package main

import (
	"Allfather/rest/server"
	"Allfather/site"
)

const port = "8080"

func main() {
	serverInstance := server.New()

	server.AddEndpoint(
		serverInstance,
		site.Index{Message: "Hello Go"})

	server.Serve(serverInstance, port)
}
