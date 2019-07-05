package basetypes

import (
	"github.com/shopspring/decimal"
	"time"
)

/*ITrade ...*/
type ITrade interface {
	Data

	Price() decimal.Decimal
	Volume() decimal.Decimal
	Time() time.Time
	StockCode() string
	Aggressor() string
	IsAfterMarket() bool
	NumberOfStocks() int
}
