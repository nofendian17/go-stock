package usecase

import (
	"context"
	"fmt"
	"go-stock/internal/entity"
	"go-stock/internal/infrastructure/idx"
	"go-stock/internal/repository"
	"go-stock/internal/shared/helper"
)

type StockSummaryUseCase interface {
	UpdateSummaries(ctx context.Context, date string) error
	FindSummaries(ctx context.Context, stockCode string, startDate, endDate string) ([]entity.StockSummary, error)
}

type stockSummaryUseCase struct {
	stockSummaryRepository repository.StockSummaryRepository
	idxClient              idx.IdxClient
}

func NewStockSummaryUseCase(idxClient idx.IdxClient, stockSummaryRepository repository.StockSummaryRepository) StockSummaryUseCase {
	return &stockSummaryUseCase{
		stockSummaryRepository: stockSummaryRepository,
		idxClient:              idxClient,
	}
}

func (b *stockSummaryUseCase) FindSummaries(ctx context.Context, code string, startDate, endDate string) ([]entity.StockSummary, error) {
	return b.stockSummaryRepository.Find(ctx, code, startDate, endDate)
}

func (b *stockSummaryUseCase) UpdateSummaries(ctx context.Context, date string) error {
	list, err := b.idxClient.GetStockSummaryList(ctx, date)
	if err != nil {
		return err
	}

	var stockSummaries []entity.StockSummary
	for _, stockSummary := range list.StockSummaryListData {
		stockSummaries = append(stockSummaries, entity.StockSummary{
			IDStockSummary:      stockSummary.IDStockSummary,
			Date:                helper.StringToDate(stockSummary.Date),
			StockCode:           stockSummary.StockCode,
			StockName:           stockSummary.StockName,
			Remarks:             stockSummary.Remarks,
			Previous:            stockSummary.Previous,
			OpenPrice:           stockSummary.OpenPrice,
			FirstTrade:          stockSummary.FirstTrade,
			High:                stockSummary.High,
			Low:                 stockSummary.Low,
			Close:               stockSummary.Close,
			Change:              stockSummary.Change,
			Volume:              stockSummary.Volume,
			Value:               stockSummary.Value,
			Frequency:           stockSummary.Frequency,
			IndexIndividual:     stockSummary.IndexIndividual,
			Offer:               stockSummary.Offer,
			OfferVolume:         stockSummary.OfferVolume,
			Bid:                 stockSummary.Bid,
			BidVolume:           stockSummary.BidVolume,
			ListedShares:        stockSummary.ListedShares,
			TradebleShares:      stockSummary.TradebleShares,
			WeightForIndex:      stockSummary.WeightForIndex,
			ForeignSell:         stockSummary.ForeignSell,
			ForeignBuy:          stockSummary.ForeignBuy,
			DelistingDate:       stockSummary.DelistingDate,
			NonRegularVolume:    stockSummary.NonRegularVolume,
			NonRegularValue:     stockSummary.NonRegularValue,
			NonRegularFrequency: stockSummary.NonRegularFrequency,
			Persen:              stockSummary.Persen,
			Percentage:          stockSummary.Percentage,
		})
	}

	if len(stockSummaries) == 0 {
		return nil // no broker data to update
	}

	if err := b.stockSummaryRepository.BulkUpsert(ctx, stockSummaries); err != nil {
		return fmt.Errorf("bulk upsert failed: %w", err)
	}

	return nil
}
