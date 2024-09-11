package adapters

import "fmt"

// - Ну типа доделать реализацию коннекта с бд, ну по идее и после этого можно будет доделать eventStore; - СУКА НАОБОРОТ ДОЛБАЕБ; (- я сам с собой если чо)
type PostgresAdapter struct{}

func NewPostgresAdapter() *PostgresAdapter {
	return &PostgresAdapter{}
}

func (pr *PostgresAdapter) Get(aggregateId string) ([]interface{}, error) {

	return nil, fmt.Errorf("not impl")
}

func (pr *PostgresAdapter) Save(data interface{}) error {
	return fmt.Errorf("not impl")
}

func (pr *PostgresAdapter) GetByAggregateId(aggregateId string) ([]interface{}, error) {
	return nil, fmt.Errorf("not impl")
}
