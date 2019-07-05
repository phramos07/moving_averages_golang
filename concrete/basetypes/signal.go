package basetypes

import (
	"fmt"
	"github.com/google/uuid"
	abstracttype "ma/abstract/basetypes"
)

/*Signal ...*/
type Signal struct {
	abstracttype.ISignal

	signal abstracttype.SignalType
	ID     uuid.UUID
}

/*NewSignal ...*/
func NewSignal(signal abstracttype.SignalType) abstracttype.ISignal {
	return &Signal{
		signal: signal,
		ID:     uuid.New(),
	}
}

/*Get ...*/
func (s *Signal) Get() abstracttype.SignalType {
	return s.signal
}

/*GUID ...*/
func (s *Signal) GUID() string {
	return s.ID.String()
}

/*String ...*/
func (s *Signal) String() string {
	return fmt.Sprintf("%s", s.Get())
}
