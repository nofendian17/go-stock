package model

import "time"

type StockSummaryRequest struct {
	StockCode string `json:"stock_code,omitempty" validate:"omitempty,len=4"`
	StartDate string `json:"startDate,omitempty" validate:"required,datetime=2006-01-02"`
	EndDate   string `json:"endDate,omitempty" validate:"required,datetime=2006-01-02"`
}

type StockSummaryResponse struct {
	IDStockSummary      int         `json:"id_stock_summary"`
	Date                time.Time   `json:"date"`
	StockCode           string      `json:"stock_code"`
	StockName           string      `json:"stock_name"`
	Remarks             string      `json:"remarks"`
	Previous            float64     `json:"previous"`
	OpenPrice           float64     `json:"open_price"`
	FirstTrade          float64     `json:"first_trade"`
	High                float64     `json:"high"`
	Low                 float64     `json:"low"`
	Close               float64     `json:"close"`
	Change              float64     `json:"change"`
	Volume              float64     `json:"volume"`
	Value               float64     `json:"value"`
	Frequency           float64     `json:"frequency"`
	IndexIndividual     float64     `json:"index_individual"`
	Offer               float64     `json:"offer"`
	OfferVolume         float64     `json:"offer_volume"`
	Bid                 float64     `json:"bid"`
	BidVolume           float64     `json:"bid_volume"`
	ListedShares        float64     `json:"listed_shares"`
	TradebleShares      float64     `json:"tradeble_shares"`
	WeightForIndex      float64     `json:"weight_for_index"`
	ForeignSell         float64     `json:"foreign_sell"`
	ForeignBuy          float64     `json:"foreign_buy"`
	DelistingDate       string      `json:"delisting_date"`
	NonRegularVolume    float64     `json:"non_regular_volume"`
	NonRegularValue     float64     `json:"non_regular_value"`
	NonRegularFrequency float64     `json:"non_regular_frequency"`
	Persen              interface{} `json:"persen"`
	Percentage          interface{} `json:"percentage"`
}
