package interfaces

type EventSender interface {
	SendEvent(e interface{}, eventType, source string) error
}
