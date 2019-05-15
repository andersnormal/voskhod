package proto

func NewReserve() *Event {
	return &Event{
		Event: &Event_Reserve{
			Reserve: &Reserve{
				Memory: 1,
			},
		},
	}
}
