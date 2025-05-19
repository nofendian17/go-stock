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
	"time"
)

type stockSummaryRepository struct {
	cfg         config.Config
	mongoClient MongoClient
	collection  string
}

func NewStockSummaryRepository(cfg config.Config, mongoClient MongoClient, collection string) repository.StockSummaryRepository {
	return &stockSummaryRepository{
		cfg:         cfg,
		mongoClient: mongoClient,
		collection:  collection,
	}
}

func (r *stockSummaryRepository) BulkUpsert(ctx context.Context, stockSummaries []entity.StockSummary) error {
	collection := r.mongoClient.GetClient().
		Database(r.cfg.GetMongo().Database).
		Collection(r.collection)

	var models []mongo.WriteModel
	for _, stockSummary := range stockSummaries {
		filter := bson.M{
			"stock_code": stockSummary.StockCode,
			"date":       stockSummary.Date,
		}
		update := bson.M{"$set": stockSummary}

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

func (r *stockSummaryRepository) Find(ctx context.Context, stockCode string, startDate, endDate string) ([]entity.StockSummary, error) {
	collection := r.mongoClient.GetClient().
		Database(r.cfg.GetMongo().Database).
		Collection(r.collection)

	filter := bson.M{}
	dateFilter := bson.M{}
	if startDate != "" {
		start, err := time.Parse("2006-01-02", startDate)
		if err == nil {
			dateFilter["$gte"] = start
		}
	}
	if endDate != "" {
		end, err := time.Parse("2006-01-02", endDate)
		if err == nil {
			dateFilter["$lte"] = end
		}
	}
	if len(dateFilter) > 0 {
		filter["date"] = dateFilter
	}

	if stockCode != "" {
		filter["stock_code"] = stockCode
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("find failed: %w", err)
	}
	defer cursor.Close(ctx)

	var results []entity.StockSummary
	if err := cursor.All(ctx, &results); err != nil {
		return nil, fmt.Errorf("decode failed: %w", err)
	}

	return results, nil
}
