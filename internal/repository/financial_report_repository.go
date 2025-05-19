package repository

import (
	"context"
	"go-stock/internal/entity"
)

type FinancialReportRepository interface {
	BulkUpsert(ctx context.Context, brokers []entity.FinancialReport) error
	Find(ctx context.Context, stockCode, reportPeriod, reportYear string) (*entity.FinancialReport, error)
}
