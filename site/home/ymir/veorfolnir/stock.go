package veorfolnir

type Stocks struct {
	Symbol string
	Price int
	Exponent int
}

type StockQuery interface {
	GetLatest() []Stocks
}
