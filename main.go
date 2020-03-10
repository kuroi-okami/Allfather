package main

import (
	"Allfather/rest/server"
	"Allfather/site/home"
	"Allfather/site/home/ymir/api/v1"
	"Allfather/site/home/ymir/veorfolnir"
	"fmt"
	"time"
)

const port = "8080"

var timer *time.Timer

func main() {
	veorfolnir.FetchStocks()
	timer = time.AfterFunc(time.Minute, func() {
		fmt.Println("Timer has expired, fetching stocks")
		veorfolnir.FetchStocks()
		timer.Reset(time.Minute)
	})
	defer timer.Stop()

	serverInstance := server.New()

	server.AddEndpoint(
		serverInstance,
		home.Index{Message: "Hello Go"})

	server.AddEndpoint(
		serverInstance,
		ymir.Index)


	server.Serve(serverInstance, port)
}

