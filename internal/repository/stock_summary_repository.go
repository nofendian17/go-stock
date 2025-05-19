package repository

import (
	"context"
	"go-stock/internal/entity"
)

type StockSummaryRepository interface {
	BulkUpsert(ctx context.Context, summaries []entity.StockSummary) error
	Find(ctx context.Context, code string, startDate, endDate string) ([]entity.StockSummary, error)
}
