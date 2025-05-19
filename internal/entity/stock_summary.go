package entity

import "time"

type StockSummary struct {
	IDStockSummary      int         `bson:"id_stock_summary"`
	Date                time.Time   `bson:"date"`
	StockCode           string      `bson:"stock_code"`
	StockName           string      `bson:"stock_name"`
	Remarks             string      `bson:"remarks"`
	Previous            float64     `bson:"previous"`
	OpenPrice           float64     `bson:"open_price"`
	FirstTrade          float64     `bson:"first_trade"`
	High                float64     `bson:"high"`
	Low                 float64     `bson:"low"`
	Close               float64     `bson:"close"`
	Change              float64     `bson:"change"`
	Volume              float64     `bson:"volume"`
	Value               float64     `bson:"value"`
	Frequency           float64     `bson:"frequency"`
	IndexIndividual     float64     `bson:"index_individual"`
	Offer               float64     `bson:"offer"`
	OfferVolume         float64     `bson:"offer_volume"`
	Bid                 float64     `bson:"bid"`
	BidVolume           float64     `bson:"bid_volume"`
	ListedShares        float64     `bson:"listed_shares"`
	TradebleShares      float64     `bson:"tradeble_shares"`
	WeightForIndex      float64     `bson:"weight_for_index"`
	ForeignSell         float64     `bson:"foreign_sell"`
	ForeignBuy          float64     `bson:"foreign_buy"`
	DelistingDate       string      `bson:"delisting_date"`
	NonRegularVolume    float64     `bson:"non_regular_volume"`
	NonRegularValue     float64     `bson:"non_regular_value"`
	NonRegularFrequency float64     `bson:"non_regular_frequency"`
	Persen              interface{} `bson:"persen"`
	Percentage          interface{} `bson:"percentage"`
}
