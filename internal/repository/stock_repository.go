package repository

import (
	"context"
	"go-stock/internal/entity"
)

type StockRepository interface {
	BulkUpsert(ctx context.Context, stocks []entity.Stock) error
	All(ctx context.Context) ([]entity.Stock, error)
	FindOne(ctx context.Context, code string) (*entity.Stock, error)
	FindWithPagination(ctx context.Context, limit, offset int64) ([]entity.Stock, int64, error)
	Search(ctx context.Context, query string) ([]entity.Stock, error)
}
