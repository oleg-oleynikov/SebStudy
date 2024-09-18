package server

import (
	"context"
	"strings"
)

func (s *server) initMongoDBCollections(ctx context.Context) {
	err := s.mongoClient.Database(s.cfg.Mongo.Db).CreateCollection(ctx, s.cfg.MongoCollections.Resumes)
	if err != nil {
		if !strings.Contains(err.Error(), "already exists") {
			s.log.Debugf("(CreateCollection) err: %v", err)
		}
	}

}
