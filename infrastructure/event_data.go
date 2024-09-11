package infrastructure

type EventData struct {
	Event         interface{}   `json:"event"`
	EventMetadata EventMetadata `json:"event_metadata"`
}
