package veorfolnir

import (
	"Allfather/utility/db/aurora"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456789"
	dbname   = "veorfolnir"
)

var psqlInfo = fmt.Sprintf(
	"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	host,
	port,
	user,
	password,
	dbname)

type Veorfolnir interface {
	initDb()
	UpdateTable()
}

type veorfolnirImpl struct {
	stocks []Stocks
	time time.Time
}

func (f veorfolnirImpl) initDb() {
	db, err := sql.Open("postgres", aurora.MakeAuroraConnection(aurora.Endpoint, aurora.DbName))
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS stocks(\n    name varchar(20) UNIQUE NOT NULL,\n    price INT NOT NULL,\n    exponent INT NOT NULL,\n    date TIMESTAMP NOT NULL\n);")
	if err != nil {
		panic(err)
	}
}
func (f veorfolnirImpl) UpdateTable(){
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	for _, stock := range f.stocks {
		_, err = db.Exec(
			"INSERT INTO stocks(\n    name, price, exponent, date) \nVALUES ($1,$2,$3,$4)\nON CONFLICT (name) DO UPDATE SET price=EXCLUDED.price, exponent=EXCLUDED.exponent, date=EXCLUDED.date;",
			stock.Symbol, stock.Price, stock.Exponent, f.time)
		if err != nil {
			panic(err)
		}
	}
}

func MakeVeorfolnirConnection(stocks []Stocks) Veorfolnir{
	var iface Veorfolnir = veorfolnirImpl {
		stocks,
		time.Now(),
	}

	// Ensure table is ready to go
	iface.initDb()
	return iface
}