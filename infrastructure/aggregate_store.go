package infrastructure

type AggregateStore interface {
	Save(a AggregateRoot, m CommandMetadata) error
	Load(id string, a AggregateRoot) error
}
