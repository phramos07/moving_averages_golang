package basetypes

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"ma/abstract/basetypes"
	"time"
)

/*Trade ...*/
type Trade struct {
	price  decimal.Decimal
	volume decimal.Decimal

	time time.Time

	identifier string
	stockCode  string
	aggressor  string

	isAfterMarket bool

	numberOfStocks int

	ID uuid.UUID
}

/*NewTrade ...*/
func NewTrade(price, volume decimal.Decimal, time time.Time,
	identifier, stockCode, aggressor string, isAfterMarket bool,
	numberOfStocks int) basetypes.ITrade {
	return &Trade{
		price:          price,
		volume:         volume,
		time:           time,
		identifier:     identifier,
		stockCode:      stockCode,
		aggressor:      aggressor,
		isAfterMarket:  isAfterMarket,
		numberOfStocks: numberOfStocks,
		ID:             uuid.New(),
	}
}

/*Price ...*/
func (t *Trade) Price() decimal.Decimal {
	return t.price
}

/*Volume ...*/
func (t *Trade) Volume() decimal.Decimal {
	return t.volume
}

/*Time ...*/
func (t *Trade) Time() time.Time {
	return t.time
}

/*StockCode ...*/
func (t *Trade) StockCode() string {
	return t.stockCode
}

/*Aggressor ...*/
func (t *Trade) Aggressor() string {
	return t.aggressor
}

/*IsAfterMarket ...*/
func (t *Trade) IsAfterMarket() bool {
	return t.isAfterMarket
}

/*NumberOfStocks ...*/
func (t *Trade) NumberOfStocks() int {
	return t.numberOfStocks
}

/*GUID ...*/
func (t *Trade) GUID() string {
	return t.ID.String()
}

/*String ...*/
func (t *Trade) String() string {
	return fmt.Sprintf(
		"%s [TRADE] %s PRICE: %s VOLUME: %s AFTERMKT: %t AGG: %s NSTOCKS: %d ",
		t.time.String(),
		t.stockCode,
		t.price.String(),
		t.volume.String(),
		t.isAfterMarket,
		t.aggressor,
		t.numberOfStocks,
	)
}
