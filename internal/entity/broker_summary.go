package entity

import "time"

type BrokerSummary struct {
	StockCode string              `bson:"stock_code"`
	StartDate time.Time           `bson:"start_date"`
	EndDate   time.Time           `bson:"end_date"`
	Buyers    []BrokerSummaryData `bson:"buyers"`
	Sellers   []BrokerSummaryData `bson:"sellers"`
	Summary   Summary             `bson:"summary"`
}

type BrokerSummaryData struct {
	BrokerCode string  `bson:"broker_code"`
	Lot        float64 `bson:"lot"`
	Val        string  `bson:"val"`
	Avg        float64 `bson:"avg"`
}

type Summary struct {
	TotalVal      string  `bson:"total_val"`
	ForeignNetVal string  `bson:"foreign_net_val"`
	TotalLot      float64 `bson:"total_lot"`
	Avg           float64 `bson:"avg"`
}
