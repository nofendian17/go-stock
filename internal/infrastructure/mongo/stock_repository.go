package mongo

import (
	"context"
	"errors"
	"fmt"
	"go-stock/internal/config"
	"go-stock/internal/entity"
	"go-stock/internal/repository"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type stockRepository struct {
	cfg         config.Config
	mongoClient MongoClient
	collection  string
}

func NewStockRepository(cfg config.Config, mongoClient MongoClient, collection string) repository.StockRepository {
	return &stockRepository{
		cfg:         cfg,
		mongoClient: mongoClient,
		collection:  collection,
	}
}

func (r *stockRepository) BulkUpsert(ctx context.Context, stocks []entity.Stock) error {
	collection := r.mongoClient.GetClient().
		Database(r.cfg.GetMongo().Database).
		Collection(r.collection)

	var models []mongo.WriteModel
	for _, stock := range stocks {
		filter := bson.M{
			"stock_code": stock.StockCode,
		}
		update := bson.M{"$set": stock}

		model := mongo.NewUpdateOneModel().
			SetFilter(filter).
			SetUpdate(update).
			SetUpsert(true)

		models = append(models, model)
	}

	if len(models) == 0 {
		return nil // No stocks to upsert
	}

	opts := options.BulkWrite().SetOrdered(false)
	_, err := collection.BulkWrite(ctx, models, opts)
	if err != nil {
		return fmt.Errorf("bulk upsert failed: %w", err)
	}

	return nil
}

func (r *stockRepository) All(ctx context.Context) ([]entity.Stock, error) {
	collection := r.mongoClient.GetClient().
		Database(r.cfg.GetMongo().Database).
		Collection(r.collection)

	cursor, err := collection.Find(
		ctx,
		bson.D{}, // Use bson.D over bson.M for consistent ordering
		options.Find().SetSort(bson.D{{Key: "stock_code", Value: 1}}),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to find stocks: %w", err)
	}
	defer cursor.Close(ctx)

	var stocks []entity.Stock
	if err := cursor.All(ctx, &stocks); err != nil {
		return nil, fmt.Errorf("failed to decode stocks: %w", err)
	}

	return stocks, nil
}

func (r *stockRepository) FindOne(ctx context.Context, code string) (*entity.Stock, error) {
	collection := r.mongoClient.GetClient().
		Database(r.cfg.GetMongo().Database).
		Collection(r.collection)

	filter := bson.M{"stock_code": code}
	var stock entity.Stock
	err := collection.FindOne(ctx, filter).Decode(&stock)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to find stock by code: %w", err)
	}
	return &stock, nil
}
