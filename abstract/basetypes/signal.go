package basetypes

/*SignalType ...*/
type SignalType string

/*ISignal ...*/
type ISignal interface {
	Data

	Get() SignalType
}

//
const (
	HALT SignalType = "HALT"
)
