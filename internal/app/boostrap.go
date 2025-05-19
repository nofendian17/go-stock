package app

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"go-stock/internal/config"
	"go-stock/internal/delivery/http/handler"
	"go-stock/internal/infrastructure/idx"
	"go-stock/internal/infrastructure/indopremier"
	"go-stock/internal/infrastructure/mongo"
	"go-stock/internal/repository"
	"go-stock/internal/usecase"
	"net/http"
	"time"
)

type Bootstrap interface {
	GetConfig() config.Config
	GetInfrastructure() Infrastructure
	GetRepository() Repository
	GetUsecase() Usecase
	GetHandler() Handler
}

type Infrastructure struct {
	mongoClient       mongo.MongoClient
	idxClient         idx.IdxClient
	indopremierClient indopremier.IndopremierClient
}

type Repository struct {
	StockRepository           repository.StockRepository
	StockSummaryRepository    repository.StockSummaryRepository
	BrokerRepository          repository.BrokerRepository
	FinancialReportRepository repository.FinancialReportRepository
}

type Usecase struct {
	StockUsecase           usecase.StockUseCase
	StockSummaryUsecase    usecase.StockSummaryUseCase
	BrokerUsecase          usecase.BrokerUseCase
	FinancialReportUseCase usecase.FinancialReportUseCase
	BrokerSummaryUseCase   usecase.BrokerSummaryUseCase
}

type Handler struct {
	HealthHandler          handler.HealthHandler
	StockHandler           handler.StockHandler
	StockSummaryHandler    handler.StockSummaryHandler
	BrokerHandler          handler.BrokerHandler
	BrokerSummaryHandler   handler.BrokerSummaryHandler
	FinancialReportHandler handler.FinancialReportHandler
}

type bootstrap struct {
	cfg            config.Config
	infrastructure Infrastructure
	repository     Repository
	usecase        Usecase
	handler        Handler
}

func (b *bootstrap) GetConfig() config.Config {
	return b.cfg
}

func (b *bootstrap) GetInfrastructure() Infrastructure {
	return b.infrastructure
}

func (b *bootstrap) GetRepository() Repository {
	return b.repository
}

func (b *bootstrap) GetUsecase() Usecase {
	return b.usecase
}

func (b *bootstrap) GetHandler() Handler {
	return b.handler
}

func NewBootstrap(cfg config.Config) (Bootstrap, error) {
	mongoConfig := mongo.Config{
		Dsn: cfg.GetMongo().DSN,
	}

	mongoClient, err := mongo.NewClient(&mongoConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize MongoDB client: %w", err)
	}

	httpClient := &http.Client{}
	idxClient := idx.NewIdxClient(idx.Config{
		BaseURL: cfg.GetService().IDXService.BaseURL,
		Delay:   time.Duration(cfg.GetService().IDXService.Delay) * time.Millisecond,
		Path: idx.Path{
			StockList:        cfg.GetService().IDXService.Path.StockList,
			StockSummaryList: cfg.GetService().IDXService.Path.StockSummaryList,
			BrokerList:       cfg.GetService().IDXService.Path.BrokerList,
			CompanyProfile:   cfg.GetService().IDXService.Path.CompanyProfile,
			FinancialReport:  cfg.GetService().IDXService.Path.FinancialReport,
		},
	}, httpClient)

	indopremierClient := indopremier.NewIndopremierClient(indopremier.Config{
		BaseURL: cfg.GetService().IndoPremierService.BaseURL,
		Delay:   time.Duration(cfg.GetService().IDXService.Delay) * time.Millisecond,
		Path: indopremier.Path{
			BrokerSummary: cfg.GetService().IndoPremierService.Path.BrokerSummary,
		},
	}, httpClient)

	stockRepository := mongo.NewStockRepository(cfg, mongoClient, "stocks")
	stockUsecase := usecase.NewStockUsecase(cfg, idxClient, stockRepository)

	stockSummaryRepository := mongo.NewStockSummaryRepository(cfg, mongoClient, "stock_summaries")
	stockSummaryUsecase := usecase.NewStockSummaryUseCase(idxClient, stockSummaryRepository)

	brokerRepository := mongo.NewBrokerRepository(cfg, mongoClient, "brokers")
	brokerUsecase := usecase.NewBrokerUseCase(idxClient, brokerRepository)

	financialReportRepository := mongo.NewFinancialReportRepository(cfg, mongoClient, "financial_reports")
	financialReportUsecase := usecase.NewFinancialReportUseCase(cfg, idxClient, financialReportRepository)

	brokerSummaryUsecase := usecase.NewBrokerSummaryUseCase(indopremierClient)

	validate := validator.New()

	healthHandler := handler.NewHealthHandler()
	stockHandler := handler.NewStockHandler(stockUsecase, validate)
	stockSummaryHandler := handler.NewStockSummaryHandler(stockSummaryUsecase, validate)
	brokerHandler := handler.NewBrokerHandler(brokerUsecase, validate)
	brokerSummaryHandler := handler.NewBrokerSummaryHandler(brokerSummaryUsecase, validate)
	financialReportHandler := handler.NewFinancialReportHandler(financialReportUsecase, validate)

	return &bootstrap{
		cfg: cfg,
		infrastructure: Infrastructure{
			mongoClient:       mongoClient,
			idxClient:         idxClient,
			indopremierClient: indopremierClient,
		},
		repository: Repository{
			StockRepository:           stockRepository,
			StockSummaryRepository:    stockSummaryRepository,
			BrokerRepository:          brokerRepository,
			FinancialReportRepository: financialReportRepository,
		},
		usecase: Usecase{
			StockUsecase:           stockUsecase,
			StockSummaryUsecase:    stockSummaryUsecase,
			BrokerUsecase:          brokerUsecase,
			FinancialReportUseCase: financialReportUsecase,
			BrokerSummaryUseCase:   brokerSummaryUsecase,
		},
		handler: Handler{
			HealthHandler:          healthHandler,
			StockHandler:           stockHandler,
			StockSummaryHandler:    stockSummaryHandler,
			BrokerHandler:          brokerHandler,
			BrokerSummaryHandler:   brokerSummaryHandler,
			FinancialReportHandler: financialReportHandler,
		},
	}, nil
}
