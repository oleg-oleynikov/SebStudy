package db_ports

type ReadModel interface {
	Get(aggregateId int, collections string) ([]interface{}, error)
	Save(data interface{}) error
}
