package db_ports

type WriteModel interface {
	Get(aggregateId int, collections string) ([]interface{}, error)
	Save(data interface{}) error
}
