package ports

type CeEventSender interface {
	SendEvent(e interface{}, eventType, source string) error
}
