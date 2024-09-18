package repository

import (
	"SebStudy/config"
	"SebStudy/internal/domain/resume/models"
	"SebStudy/logger"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoRepository struct {
	ResumeMongoRepository

	log logger.Logger
	cfg *config.Config
	db  *mongo.Client
}

func NewMongoRepository(log logger.Logger, cfg *config.Config, db *mongo.Client) *mongoRepository {
	return &mongoRepository{
		log: log,
		cfg: cfg,
		db:  db,
	}
}

func (m *mongoRepository) Insert(ctx context.Context, resume *models.ResumeProjection) error {
	_, err := m.getResumesCollection().InsertOne(ctx, resume, &options.InsertOneOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (m *mongoRepository) Update(ctx context.Context, resume *models.ResumeProjection) error {
	return nil
}

func (m *mongoRepository) getResumesCollection() *mongo.Collection {
	return m.db.Database(m.cfg.Mongo.Db).Collection(m.cfg.MongoCollections.Resumes)
}
