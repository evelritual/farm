package event

type Event string

const (
	EventDownDepress  Event = "EventDownDepress"
	EventDownPress    Event = "EventDownPress"
	EventLeftDepress  Event = "EventLeftDepress"
	EventLeftPress    Event = "EventLeftPress"
	EventRightDepress Event = "EventRightDepress"
	EventRightPress   Event = "EventRightPress"
	EventUpDepress    Event = "EventUpDepress"
	EventUpPress      Event = "EventUpPress"
)
