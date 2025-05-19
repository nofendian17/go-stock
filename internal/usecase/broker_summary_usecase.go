package usecase

import (
	"context"
	"go-stock/internal/entity"
	"go-stock/internal/infrastructure/indopremier"
)

type BrokerSummaryUseCase interface {
	Find(ctx context.Context, stockCode, startDate, endDate, investorType, board string) (*entity.BrokerSummary, error)
}

type brokerSummaryUseCase struct {
	indopremierClient indopremier.IndopremierClient
}

func NewBrokerSummaryUseCase(indopremierClient indopremier.IndopremierClient) BrokerSummaryUseCase {
	return &brokerSummaryUseCase{
		indopremierClient: indopremierClient,
	}
}

func (b *brokerSummaryUseCase) Find(ctx context.Context, stockCode, startDate, endDate, investorType, board string) (*entity.BrokerSummary, error) {
	result, err := b.indopremierClient.GetBrokerSummary(ctx, stockCode, startDate, endDate, investorType, board)
	if err != nil {
		return nil, err
	}

	buyers := make([]entity.BrokerSummaryData, 0, len(result.Buyers))
	for _, buyer := range result.Buyers {
		buyers = append(buyers, entity.BrokerSummaryData{
			BrokerCode: buyer.BrokerCode,
			Lot:        buyer.Lot,
			Val:        buyer.Val,
			Avg:        buyer.Avg,
		})
	}

	sellers := make([]entity.BrokerSummaryData, 0, len(result.Sellers))
	for _, seller := range result.Sellers {
		sellers = append(sellers, entity.BrokerSummaryData{
			BrokerCode: seller.BrokerCode,
			Lot:        seller.Lot,
			Val:        seller.Val,
			Avg:        seller.Avg,
		})
	}
	return &entity.BrokerSummary{
		StockCode: result.StockCode,
		StartDate: result.StartDate,
		EndDate:   result.EndDate,
		Buyers:    buyers,
		Sellers:   sellers,
		Summary: entity.Summary{
			TotalVal:      result.Summary.TotalVal,
			ForeignNetVal: result.Summary.ForeignNetVal,
			TotalLot:      result.Summary.TotalLot,
			Avg:           result.Summary.Avg,
		},
	}, nil

}
