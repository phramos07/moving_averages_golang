package rules

// import "io/ioutil"
// import "encoding/json"
import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"log"
	abstracttype "ma/abstract/basetypes"
	"ma/abstract/rules"
	concretetype "ma/concrete/basetypes"
	"strconv"
	"time"
)

/*JSONTradeStreamer ...*/
type JSONTradeStreamer struct {
	rules.SourceStreamer

	filepath  string
	stockCode string
}

/*NewJSONTradeStreamer ...*/
func NewJSONTradeStreamer(filepath string, stockCode string) rules.SourceStreamer {
	log.Printf("Opening a new JSON Trade Streamer.")

	return &JSONTradeStreamer{
		filepath:  filepath,
		stockCode: stockCode,
	}
}

/*Stream ...*/
func (ts *JSONTradeStreamer) Stream(trades chan<- interface{}) {
	tradesFromJSON, err := fetchTradesFromJSON(ts.filepath)

	if err != nil {
		log.Panicf("Could not open input %s file.", ts.filepath)
	}

	// Send all trades through channel (as long as there is a receiver)
	for _, trade := range tradesFromJSON {
		trades <- trade
		log.Printf("Streaming new trade: %s", trade.String())
	}

	// In the end, send signal to halt robot.
	haltSignal := concretetype.NewSignal(abstracttype.HALT)
	trades <- haltSignal
}

func fetchTradesFromJSON(jsonpath string) (trades []abstracttype.Data, err error) {
	content, err := ioutil.ReadFile(jsonpath)
	if err != nil {
		panic("Json Candle File not found")
	}

	var data []map[string]string
	err = json.Unmarshal(content, &data)

	for _, d := range data {
		time := parseTime(d["time"])
		price, _ := decimal.NewFromString(d["price"])
		volume, _ := decimal.NewFromString(d["volume"])
		isAfterMarket := false
		if d["isAfterMarket"] == "True" {
			isAfterMarket = true
		}
		nstocks, _ := strconv.Atoi(d["nstocks"])

		trade := concretetype.NewTrade(price, volume, time, d["identifier"], d["stock"], d["aggressor"], isAfterMarket, nstocks)
		trades = append(trades, trade)
	}

	return trades, err
}

func parseTime(str string) (tradeTime time.Time) {
	var y, m, d, h, min, sec int
	_, parseErr := fmt.Sscanf(str, "%d-%d-%d %d:%d:%d", &y, &m, &d, &h, &min, &sec)
	if parseErr != nil {
		log.Fatal(parseErr)
	}
	tradeTime = time.Date(y, time.Month(m), d, h, min, sec, 0, time.UTC)
	return
}
