package rules

import (
	"log"
	abstracttype "ma/abstract/basetypes"
	"ma/abstract/rules"
	"reflect"
	// "time"
)

/*BaseStreamer ...*/
type BaseStreamer struct {
	rules.Streamer

	transformations		[]func(in interface{}) interface{}
}

/*NewStreamer ...*/
func NewStreamer() rules.Streamer {
	return &BaseStreamer{}
}

/*Stream ...*/
func (bs *BaseStreamer) Stream(
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

func (bs *BaseStreamer) Register(transform func (in interface{}) interface{}, T reflect.Type) {

}