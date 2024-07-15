package db_ports

type WriteModel interface {
	Get(aggregateId string) ([]interface{}, error)
	Save(data interface{}) error
	GetByAggregateId(aggregateId string) ([]interface{}, error)
}
