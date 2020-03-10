package veorfolnir

const stockApi = "https://financialmodelingprep.com/api/v3/stock/real-time-price"

func FetchStocks() {
	stockQuery := StockQueryImpl{
		Api: stockApi,
	}

	stocks := stockQuery.GetLatest()

	dbCon := MakeVeorfolnirConnection(stocks)
	dbCon.UpdateTable()
}
