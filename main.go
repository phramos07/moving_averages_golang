package main

import (
	"github.com/shopspring/decimal"
	"log"
	abstracttype "ma/abstract/basetypes"
	abstractrules "ma/abstract/rules"
	concreterules "ma/concrete/rules"
	"sync"
)

func main() {
	// Start modules
	tradeStreamer := concreterules.NewJSONTradeStreamer("data/trades.json", "WDOZ18")
	movingAverageStreamer := concreterules.NewMovingAverageStreamer(abstractrules.SIMPLE, 9)

	// Declare channels
	var trades chan interface{}
	var movingAverage chan interface{}

	// Initialize channels
	trades = make(chan interface{})
	movingAverage = make(chan interface{})

	// Start flow

	// Keep track of routines
	var waitGroup sync.WaitGroup
	waitGroup.Add(3)

	// Routine1: TradeStreamer
	go func() {
		defer waitGroup.Done()
		tradeStreamer.Stream(trades)
	}()

	// Routine2: MovingAverageStreamer
	go func() {
		defer waitGroup.Done()
		movingAverageStreamer.Stream(trades, movingAverage)
	}()

	// Routine3: MovingAverageReceiver
	go func(movingAverage <-chan interface{}) {
		defer waitGroup.Done()
		for {
			nextItem := <-movingAverage
			switch item := nextItem.(type) {
			case decimal.Decimal:
				log.Printf("MAReceiver Received new moving average: %s", item.String())
			case abstracttype.ISignal:
				log.Printf("MAReceiver Received a ISignal: %s", item.String())
				switch item.Get() {
				case abstracttype.HALT:
					log.Printf("MAReceiver Halting.")
					return
				default:
					log.Panicf("MAReceiver Signal unrecognized. Halting module.")
				}
			default:
				log.Panicf("MAReceiver should've received ISignal or decimal.Decimal. Halting.")
			}
		}
	}(movingAverage)

	waitGroup.Wait()
	log.Printf("All routines are done now.")
	log.Printf("Reached end of execution. Disposing channels.")
	close(trades)
	close(movingAverage)
}
