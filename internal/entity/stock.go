package entity

import (
	"time"
)

type Stock struct {
	StockCode       string           `bson:"stock_code"`
	StockName       string           `bson:"stock_name"`
	Share           float64          `bson:"share"`
	ListingDate     time.Time        `bson:"listing_date"`
	Board           string           `bson:"board"`
	Profiles        []Profile        `bson:"profiles"`
	Secretaries     []Secretary      `bson:"secretaries"`
	Directors       []Director       `bson:"directors"`
	Commissioners   []Commissioner   `bson:"commissioners"`
	AuditCommittees []AuditCommittee `bson:"audit_committees"`
	Shareholders    []Shareholder    `bson:"shareholders"`
	Subsidiaries    []Subsidiary     `bson:"subsidiaries"`
	Dividends       []Dividend       `bson:"dividends"`
}

type Profile struct {
	Address      string `bson:"address"`
	BAE          string `bson:"bae"`
	Industry     string `bson:"industry"`
	SubIndustry  string `bson:"sub_industry"`
	Email        string `bson:"email"`
	Fax          string `bson:"fax"`
	MainBusiness string `bson:"main_business"`
	StockCode    string `bson:"stock_code"`
	StockName    string `bson:"stock_name"`
	TIN          string `bson:"tin"`
	Sector       string `bson:"sector"`
	SubSector    string `bson:"sub_sector"`
	ListingDate  string `bson:"listing_date"`
	Phone        string `bson:"phone"`
	Website      string `bson:"website"`
	Status       int    `bson:"status"`
	Logo         string `bson:"logo"`
}

type Secretary struct {
	Name         string `bson:"name"`
	PhoneNumber  string `bson:"phone_number"`
	Website      string `bson:"website"`
	Email        string `bson:"email"`
	Fax          string `bson:"fax"`
	MobileNumber string `bson:"mobile_number"`
}

type Director struct {
	Name         string `bson:"name"`
	Position     string `bson:"position"`
	IsAffiliated bool   `bson:"is_affiliated"`
}

type Commissioner struct {
	Name          string `bson:"name"`
	Position      string `bson:"position"`
	IsIndependent bool   `bson:"is_independent"`
}

type AuditCommittee struct {
	Name     string `bson:"name"`
	Position string `bson:"position"`
}

type Shareholder struct {
	Share        float64 `bson:"share"`
	Category     string  `bson:"category"`
	Name         string  `bson:"name"`
	IsController bool    `bson:"is_controller"`
	Percentage   float64 `bson:"percentage"`
}

type Subsidiary struct {
	BusinessFields  string  `bson:"business_fields"`
	TotalAsset      float64 `bson:"total_asset"`
	Location        string  `bson:"location"`
	Currency        string  `bson:"currency"`
	Name            string  `bson:"name"`
	Percentage      float64 `bson:"percentage"`
	Units           string  `bson:"units"`
	OperationStatus string  `bson:"operation_status"`
	CommercialYear  string  `bson:"commercial_year"`
}

type Dividend struct {
	Name                         string    `bson:"name"`
	Type                         string    `bson:"type"`
	Year                         string    `bson:"year"`
	TotalStockBonus              float64   `bson:"total_stock_bonus"`
	CashDividendPerShareCurrency string    `bson:"cash_dividend_per_share_currency"`
	CashDividendPerShare         float64   `bson:"cash_dividend_per_share"`
	CumDate                      time.Time `bson:"cum_date"`
	ExDate                       time.Time `bson:"ex_date"`
	RecordDate                   time.Time `bson:"record_date"`
	PaymentDate                  time.Time `bson:"payment_date"`
	Ratio1                       int       `bson:"ratio_1"`
	Ratio2                       int       `bson:"ratio_2"`
	CashDividendCurrency         string    `bson:"cash_dividend_currency"`
	CashDividendTotal            float64   `bson:"cash_dividend_total"`
}
