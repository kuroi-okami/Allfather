package veorfolnir

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)
type responseArr struct {
	StockList []struct {
		Symbol string `json:"symbol"`
		Value float64 `json:"price"`
	} `json:"stockList"`
}

type StockQueryImpl struct {
	Api string
}

func (s StockQueryImpl) GetLatest() []Stocks {
	resp, err := http.Get(s.Api)
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(resp.Body)
	var stocks responseArr
	err = decoder.Decode(&stocks)
	if err != nil{
		panic(err)
	}

	var genStocks []Stocks
	for _, stock := range stocks.StockList {
		valueString := strconv.FormatFloat(stock.Value, 'f', -1, 64)
		trimString := strings.Replace(valueString, ".", "", 1)
		splitString := strings.Split(valueString, ".")
		length := 0
		if len(splitString) > 1 {
			length = len(splitString[1])
		}
		price, _ := strconv.Atoi( trimString )

		thisStock := Stocks{
			Symbol:   stock.Symbol,
			Price:    price,
			Exponent: length,
		}

		genStocks = append(genStocks, thisStock)
	}

	return genStocks
}
