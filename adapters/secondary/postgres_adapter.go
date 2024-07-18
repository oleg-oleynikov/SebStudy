package secondary

// Ну типа доделать реализацию коннекта с бд, ну по идее и после этого можно будет доделать eventStore
type PostgresAdapter struct{}

func NewPostgresAdapter() *PostgresAdapter {
	return &PostgresAdapter{}
}

func (pr *PostgresAdapter) Get(aggregateId string) ([]interface{}, error) {

	return nil, nil
}

func (pr *PostgresAdapter) Save(data interface{}) error {
	return nil
}

func (pr *PostgresAdapter) GetByAggregateId(aggregateId string) ([]interface{}, error) {
	return nil, nil
}
