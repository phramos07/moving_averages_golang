package rules

import (
	"github.com/shopspring/decimal"
	"log"
	abstracttype "ma/abstract/basetypes"
	"ma/abstract/rules"
	// "time"
)

/*MovingAverageStreamer ...*/
type MovingAverageStreamer struct {
	rules.Streamer

	averageType rules.MovingAverageType
	periods     int

	period        int
	slidingWindow []decimal.Decimal
}

/*NewMovingAverageStreamer ...*/
func NewMovingAverageStreamer(
	averageType rules.MovingAverageType,
	periods int) rules.Streamer {

	MAS := &MovingAverageStreamer{
		averageType:   averageType,
		periods:       periods,
		period:        0,
		slidingWindow: make([]decimal.Decimal, 0),
	}

	return MAS
}

/*Stream ...*/
func (mas *MovingAverageStreamer) Stream(
	trades <-chan interface{},
	movingAverages chan<- interface{}) {

	for {
		select {
		case newItem := <-trades:
			switch item := newItem.(type) {
			case abstracttype.ITrade:
				log.Printf("MAStreamer received a ITrade. Price: %s", item.Price().String())
				mas.onTrade(item, movingAverages)
			case abstracttype.ISignal:
				log.Printf("MAStreamer received a ISignal: %s", item.String())
				switch item.Get() {
				case abstracttype.HALT:
					log.Printf("MAStreamer received HALT signal. Halting.")
					movingAverages <- item
					return
				default:
					log.Panicf("MAStreamer Unrecognized signal.")
				}
			default:
				log.Panicf("MAStreamer should've received an ITrade or ISignal.")
			}
		default:
			continue
		}
	}
}

func (mas *MovingAverageStreamer) onTrade(
	trade abstracttype.ITrade, movingAverage chan<- interface{}) {
	switch maType := mas.averageType; maType {
	case rules.SIMPLE:
		mas.updateSimpleMovingAverageOnTrade(trade, movingAverage)
	case rules.EXPONENTIAL:
		log.Panicf("Not implemented yet.")
	default:
		mas.updateSimpleMovingAverageOnTrade(trade, movingAverage)
	}
}

func (mas *MovingAverageStreamer) updateSimpleMovingAverageOnTrade(
	trade abstracttype.ITrade, movingAverage chan<- interface{}) {
	mas.period++
	mas.slidingWindow = append(mas.slidingWindow, trade.Price())

	if len(mas.slidingWindow) > mas.periods {
		mas.slidingWindow = mas.slidingWindow[1:]
	}

	if mas.period >= mas.periods {
		average := decimal.Avg(mas.slidingWindow[0], mas.slidingWindow[1:]...)
		movingAverage <- average
	}
}
