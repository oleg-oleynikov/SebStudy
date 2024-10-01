package server

import (
	"context"
	"fmt"
	"strings"

	"github.com/EventStore/EventStore-Client-Go/esdb"
)

func (s *server) initMongoDBCollections(ctx context.Context) {
	err := s.mongoClient.Database(s.cfg.Mongo.Db).CreateCollection(ctx, s.cfg.MongoCollections.Resumes)
	if err != nil {
		if !strings.Contains(err.Error(), "already exists") {
			s.log.Debugf("(CreateCollection) err: %v", err)
		}
	}

}

func createESDBClient(connection string) (*esdb.Client, error) {
	conn, err := esdb.ParseConnectionString(connection)
	if err != nil {
		return nil, fmt.Errorf("failed to create esdb client: %v", err)
	}

	db, err := esdb.NewClient(conn)
	if err != nil {
		return nil, err
	}

	options := esdb.ReadAllOptions{
		From: esdb.Start{},
	}

	_, err = db.ReadAll(context.Background(), options, 1)
	if err != nil {
		return nil, fmt.Errorf("failed to ping esdb")
	}

	return db, nil
}
