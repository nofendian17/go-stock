package indopremier

import "time"

type GetBrokerSummaryResponse struct {
	StockCode string
	StartDate time.Time
	EndDate   time.Time
	Buyers    []BrokerSummaryData
	Sellers   []BrokerSummaryData
	Summary   Summary
}

type BrokerSummaryData struct {
	BrokerCode string
	Lot        float64
	Val        string
	Avg        float64
}
type Summary struct {
	TotalVal      string
	ForeignNetVal string
	TotalLot      float64
	Avg           float64
}
