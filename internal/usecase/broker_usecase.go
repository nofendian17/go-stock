package usecase

import (
	"context"
	"fmt"
	"go-stock/internal/entity"
	"go-stock/internal/infrastructure/idx"
	"go-stock/internal/repository"
)

type BrokerUseCase interface {
	UpdateBroker(ctx context.Context) error
	Find(ctx context.Context, code string) ([]entity.Broker, error)
}

type brokerUseCase struct {
	brokerRepository repository.BrokerRepository
	idxClient        idx.IdxClient
}

func NewBrokerUseCase(idxClient idx.IdxClient, brokerRepository repository.BrokerRepository) BrokerUseCase {
	return &brokerUseCase{
		brokerRepository: brokerRepository,
		idxClient:        idxClient,
	}
}

func (b *brokerUseCase) UpdateBroker(ctx context.Context) error {
	list, err := b.idxClient.GetBrokerList(ctx)
	if err != nil {
		return err
	}

	var brokers []entity.Broker
	for _, broker := range list.BrokerListData {
		brokers = append(brokers, entity.Broker{
			Code:    broker.Code,
			Name:    broker.Name,
			License: broker.License,
		})
	}

	if len(brokers) == 0 {
		return nil // no broker data to update
	}

	if err := b.brokerRepository.BulkUpsert(ctx, brokers); err != nil {
		return fmt.Errorf("bulk upsert failed: %w", err)
	}

	return nil
}

func (b *brokerUseCase) Find(ctx context.Context, code string) ([]entity.Broker, error) {
	return b.brokerRepository.Find(ctx, code)
}
