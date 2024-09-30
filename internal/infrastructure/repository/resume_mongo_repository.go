package repository

import (
	"SebStudy/config"
	"SebStudy/internal/domain/resume/models"
	"SebStudy/logger"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type resumeMongoRepository struct {
	ResumeRepository

	log logger.Logger
	cfg *config.Config
	db  *mongo.Client
}

func NewResumeMongoRepository(log logger.Logger, cfg *config.Config, db *mongo.Client) *resumeMongoRepository {
	return &resumeMongoRepository{
		log: log,
		cfg: cfg,
		db:  db,
	}
}

func (m *resumeMongoRepository) Insert(ctx context.Context, resume *models.ResumeProjection) error {
	_, err := m.getResumesCollection().InsertOne(ctx, resume, &options.InsertOneOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (m *resumeMongoRepository) Update(ctx context.Context, resume *models.ResumeProjection) error {

	ops := options.FindOneAndUpdate()
	ops.SetReturnDocument(options.After)
	ops.SetUpsert(false)

	if err := m.getResumesCollection().FindOneAndUpdate(ctx, bson.M{"_id": resume.Id}, bson.M{"$set": resume}, ops).Err(); err != nil {
		return err
	}

	return nil
}

func (m *resumeMongoRepository) GetByAccountId(ctx context.Context, userId string) (*models.ResumeProjection, error) {

	var resumeProjection models.ResumeProjection
	if err := m.getResumesCollection().FindOne(ctx, bson.M{"userId": userId}).Decode(&resumeProjection); err != nil {
		return nil, err
	}

	return &resumeProjection, nil
}

func (m *resumeMongoRepository) GetById(ctx context.Context, resumeId string) (*models.ResumeProjection, error) {

	var resumeProjection models.ResumeProjection
	if err := m.getResumesCollection().FindOne(ctx, bson.M{"_id": resumeId}).Decode(&resumeProjection); err != nil {
		return nil, err
	}

	return &resumeProjection, nil
}

func (m *resumeMongoRepository) ResumeExistsByAccountId(ctx context.Context, userId string) (bool, error) {
	filter := bson.M{"userId": userId}
	count, err := m.getResumesCollection().CountDocuments(ctx, filter)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (m *resumeMongoRepository) getResumesCollection() *mongo.Collection {
	return m.db.Database(m.cfg.Mongo.Db).Collection(m.cfg.MongoCollections.Resumes)
}
