package mongo

import (
	"context"
	"fmt"
	"go-stock/internal/config"
	"go-stock/internal/entity"
	"go-stock/internal/repository"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type brokerRepository struct {
	cfg         config.Config
	mongoClient MongoClient
	collection  string
}

func NewBrokerRepository(cfg config.Config, mongoClient MongoClient, collection string) repository.BrokerRepository {
	return &brokerRepository{
		cfg:         cfg,
		mongoClient: mongoClient,
		collection:  collection,
	}
}

func (r *brokerRepository) BulkUpsert(ctx context.Context, brokers []entity.Broker) error {
	collection := r.mongoClient.GetClient().
		Database(r.cfg.GetMongo().Database).
		Collection(r.collection)

	var models []mongo.WriteModel
	for _, broker := range brokers {
		filter := bson.M{"code": broker.Code}
		update := bson.M{"$set": broker}

		model := mongo.NewUpdateOneModel().
			SetFilter(filter).
			SetUpdate(update).
			SetUpsert(true)

		models = append(models, model)
	}

	if len(models) == 0 {
		return nil // no brokers to process
	}

	opts := options.BulkWrite().SetOrdered(false)
	_, err := collection.BulkWrite(ctx, models, opts)
	if err != nil {
		return fmt.Errorf("bulk upsert failed: %w", err)
	}

	return nil
}

func (r *brokerRepository) Find(ctx context.Context, code string) ([]entity.Broker, error) {
	collection := r.mongoClient.GetClient().
		Database(r.cfg.GetMongo().Database).
		Collection(r.collection)

	filter := bson.M{}
	if code != "" {
		filter["code"] = code
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to find brokers: %w", err)
	}
	defer cursor.Close(ctx)

	var brokers []entity.Broker
	if err := cursor.All(ctx, &brokers); err != nil {
		return nil, fmt.Errorf("failed to decode brokers: %w", err)
	}

	return brokers, nil
}
