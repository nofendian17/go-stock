package http

import (
	"go-stock/internal/app"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

func RegisterRoutes(route *http.ServeMux, handler app.Handler) {

	// Define the routes
	route.HandleFunc("/healthz", handler.HealthHandler.Healthz)
	route.HandleFunc("/api/v1/stocks", handler.StockHandler.ListStock)
	route.HandleFunc("/api/v1/stock", handler.StockHandler.FindStock)
	route.HandleFunc("/api/v1/stock/summaries", handler.StockSummaryHandler.FindStockSummaries)
	route.HandleFunc("/api/v1/brokers", handler.BrokerHandler.Find)
	route.HandleFunc("/api/v1/brokers/summaries", handler.BrokerSummaryHandler.Find)
	route.HandleFunc("/api/v1/financial_report", handler.FinancialReportHandler.FindFinancialReport)

	// Swagger
	route.HandleFunc("/swagger/", httpSwagger.WrapHandler)
}
