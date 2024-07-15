package secondary

type PostgresAdapter struct{}

func NewPostgresAdapter() *PostgresAdapter {
	return &PostgresAdapter{}
}

func (pr *PostgresAdapter) Get(aggregateId string) ([]interface{}, error) {

	return nil, nil
}

func (pr *PostgresAdapter) Save(data interface{}) error {
	// log.Println("Sssssss")
	return nil
}

func (pr *PostgresAdapter) GetByAggregateId(aggregateId string) ([]interface{}, error) {
	return nil, nil
}
