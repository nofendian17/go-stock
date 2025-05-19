package repository

import (
	"context"
	"go-stock/internal/entity"
)

type BrokerRepository interface {
	BulkUpsert(ctx context.Context, brokers []entity.Broker) error
	Find(ctx context.Context, code string) ([]entity.Broker, error)
}
