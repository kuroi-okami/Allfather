package main

import (
	"Allfather/rest/server"
	"Allfather/site/home"
	"Allfather/site/home/ymir/v1"
)

const port = "8080"

func main() {
	serverInstance := server.New()

	server.AddEndpoint(
		serverInstance,
		home.Index{Message: "Hello Go"})

	server.AddEndpoint(
		serverInstance,
		ymir.Index)


	server.Serve(serverInstance, port)
}
