package model

import "time"

type BrokerSummaryRequest struct {
	StockCode       string `json:"stock_code" validate:"required"`
	StartDate       string `json:"start_date" validate:"required,datetime=2006-01-02"`
	EndDate         string `json:"end_date" validate:"required,datetime=2006-01-02"`
	InvestorType    string `json:"investor_type" validate:"required,oneof=ALL F D"`
	TransactionType string `json:"transaction_type" validate:"required,oneof=ALL RG TN NG"`
}

type BrokerSummaryResponse struct {
	StockCode string              `json:"stock_code"`
	StartDate time.Time           `json:"start_date"`
	EndDate   time.Time           `json:"end_date"`
	Buyers    []BrokerSummaryData `json:"buyers"`
	Sellers   []BrokerSummaryData `json:"sellers"`
	Summary   Summary             `json:"summary"`
}

type BrokerSummaryData struct {
	BrokerCode string  `json:"broker_code"`
	Lot        float64 `json:"lot"`
	Val        string  `json:"val"`
	Avg        float64 `json:"avg"`
}

type Summary struct {
	TotalVal      string  `json:"total_val"`
	ForeignNetVal string  `json:"foreign_net_val"`
	TotalLot      float64 `json:"total_lot"`
	Avg           float64 `json:"avg"`
}
