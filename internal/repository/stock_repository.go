package repository

import (
	"context"
	"go-stock/internal/entity"
)

type StockRepository interface {
	BulkUpsert(ctx context.Context, stocks []entity.Stock) error
	All(ctx context.Context) ([]entity.Stock, error)
	FindOne(ctx context.Context, code string) (*entity.Stock, error)
}
