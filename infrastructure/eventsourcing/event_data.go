package eventsourcing

type EventData struct {
	Data     interface{}
	Type     string
	Metadata map[string]string
}
