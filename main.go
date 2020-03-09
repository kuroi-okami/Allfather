package main

import (
	"Allfather/rest/server"
	"Allfather/site/home"
	"Allfather/site/home/ymir/api/v1"
	"Allfather/site/home/ymir/veorfolnir"
	"fmt"
)

const port = "8080"
const stockApi = "https://financialmodelingprep.com/api/v3/stock/real-time-price"

func main() {
	stocky := veorfolnir.StockQueryImpl{
		Api: stockApi,
	}

	stocks := stocky.GetLatest()
	for _, stock := range stocks {
		fmt.Printf(
			"Name: %s ", stock.Symbol)
		fmt.Printf(
			"Cost: %d ", stock.Price)
		fmt.Printf(
			"Exponant: %d ", stock.Exponent)
		fmt.Println("----------------------------")
	}

	serverInstance := server.New()

	server.AddEndpoint(
		serverInstance,
		home.Index{Message: "Hello Go"})

	server.AddEndpoint(
		serverInstance,
		ymir.Index)


	server.Serve(serverInstance, port)
}
