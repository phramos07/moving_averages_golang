package rules

import (
	"reflect"
)

/*SourceStreamer ...*/
type SourceStreamer interface {
	Stream(outgoing chan<- interface{})
}

/*Streamer ...*/
type Streamer interface {
	Stream(incoming <-chan interface{}, outgoing chan<- interface{})
	Register(transform func(in interface{}) interface{}, T reflect.Type)
}
