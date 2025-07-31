package http

import (
	"go-stock/internal/app"
	"go-stock/internal/delivery/http/middleware"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

func RegisterRoutes(mux *http.ServeMux, app app.Bootstrap) {
	cors := middleware.CORS("*")
	log := middleware.Logger()

	chain := func(h http.HandlerFunc) http.HandlerFunc {
		return middleware.Chain(h, log, cors)
	}

	// Apply middleware to routes
	mux.HandleFunc("/healthz", chain(app.GetHandler().HealthHandler.Healthz))
	mux.HandleFunc("/api/v1/stocks", chain(app.GetHandler().StockHandler.ListStock))
	mux.HandleFunc("/api/v1/stocks/search", chain(app.GetHandler().StockHandler.SearchStock))
	mux.HandleFunc("/api/v1/stock", chain(app.GetHandler().StockHandler.FindStock))
	mux.HandleFunc("/api/v1/stock/summaries", chain(app.GetHandler().StockSummaryHandler.FindStockSummaries))
	mux.HandleFunc("/api/v1/brokers", chain(app.GetHandler().BrokerHandler.Find))
	mux.HandleFunc("/api/v1/brokers/summaries", chain(app.GetHandler().BrokerSummaryHandler.Find))
	mux.HandleFunc("/api/v1/financial_report", chain(app.GetHandler().FinancialReportHandler.FindFinancialReport))

	// Swagger & Static files
	mux.HandleFunc("/swagger/", httpSwagger.WrapHandler.ServeHTTP)
	mux.Handle("/", chain(http.FileServer(http.FS(app.GetView().ViewService.GetFS())).ServeHTTP))
}
