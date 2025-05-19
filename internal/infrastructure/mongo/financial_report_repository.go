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

type financialReportRepository struct {
	cfg         config.Config
	mongoClient MongoClient
	collection  string
}

func NewFinancialReportRepository(cfg config.Config, mongoClient MongoClient, collection string) repository.FinancialReportRepository {
	return &financialReportRepository{
		cfg:         cfg,
		mongoClient: mongoClient,
		collection:  collection,
	}
}

func (r *financialReportRepository) BulkUpsert(ctx context.Context, financialReports []entity.FinancialReport) error {
	collection := r.mongoClient.GetClient().
		Database(r.cfg.GetMongo().Database).
		Collection(r.collection)

	var models []mongo.WriteModel
	for _, financialReport := range financialReports {
		filter := bson.M{
			"stock_code":    financialReport.StockCode,
			"report_period": financialReport.ReportPeriod,
			"report_year":   financialReport.ReportYear,
			"file_modified": financialReport.FileModified,
		}
		update := bson.M{"$set": financialReport}

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

func (r *financialReportRepository) Find(ctx context.Context, stockCode, reportPeriod, reportYear string) (*entity.FinancialReport, error) {
	collection := r.mongoClient.GetClient().
		Database(r.cfg.GetMongo().Database).
		Collection(r.collection)

	filter := bson.M{
		"stock_code":    stockCode,
		"report_period": reportPeriod,
		"report_year":   reportYear,
	}

	var result entity.FinancialReport
	err := collection.FindOne(ctx, filter).Decode(&result)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("find failed: %w", err)
	}
	return &result, nil
}
