package eventsourcing

type EventData struct {
	DataBytes     []byte
	MetadataBytes []byte
}

func NewEventData(dataBytes []byte, metadataBytes []byte) *EventData {
	return &EventData{
		DataBytes:     dataBytes,
		MetadataBytes: metadataBytes,
	}
}
