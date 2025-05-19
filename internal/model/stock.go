package model

import "time"

type StockRequest struct {
	StockCode string `json:"stock_code,omitempty" validate:"required,len=4"`
}

type StockResponse struct {
	Code            string           `json:"code"`
	Name            string           `json:"name"`
	Share           float64          `json:"share"`
	ListingDate     time.Time        `json:"listing_date"`
	Board           string           `json:"board"`
	Profiles        []Profile        `json:"profiles"`
	Secretaries     []Secretary      `json:"secretaries"`
	Directors       []Director       `json:"directors"`
	Commissioners   []Commissioner   `json:"commissioners"`
	AuditCommittees []AuditCommittee `json:"audit_committees"`
	Shareholders    []Shareholder    `json:"shareholders"`
	Subsidiaries    []Subsidiary     `json:"subsidiaries"`
	Dividends       []Dividend       `json:"dividends"`
}

type Profile struct {
	Address      string `json:"address"`
	BAE          string `json:"bae"`
	Industry     string `json:"industry"`
	SubIndustry  string `json:"sub_industry"`
	Email        string `json:"email"`
	Fax          string `json:"fax"`
	MainBusiness string `json:"main_business"`
	StockCode    string `json:"stock_code"`
	StockName    string `json:"stock_name"`
	TIN          string `json:"tin"`
	Sector       string `json:"sector"`
	SubSector    string `json:"sub_sector"`
	ListingDate  string `json:"listing_date"`
	Phone        string `json:"phone"`
	Website      string `json:"website"`
	Status       int    `json:"status"`
	Logo         string `json:"logo"`
}

type Secretary struct {
	Name         string `json:"name"`
	PhoneNumber  string `json:"phone_number"`
	Website      string `json:"website"`
	Email        string `json:"email"`
	Fax          string `json:"fax"`
	MobileNumber string `json:"mobile_number"`
}

type Director struct {
	Name         string `json:"name"`
	Position     string `json:"position"`
	IsAffiliated bool   `json:"is_affiliated"`
}

type Commissioner struct {
	Name          string `json:"name"`
	Position      string `json:"position"`
	IsIndependent bool   `json:"is_independent"`
}

type AuditCommittee struct {
	Name     string `json:"name"`
	Position string `json:"position"`
}

type Shareholder struct {
	Share        float64 `json:"share"`
	Category     string  `json:"category"`
	Name         string  `json:"name"`
	IsController bool    `json:"is_controller"`
	Percentage   float64 `json:"percentage"`
}

type Subsidiary struct {
	BusinessFields  string  `json:"business_fields"`
	TotalAsset      float64 `json:"total_asset"`
	Location        string  `json:"location"`
	Currency        string  `json:"currency"`
	Name            string  `json:"name"`
	Percentage      float64 `json:"percentage"`
	Units           string  `json:"units"`
	OperationStatus string  `json:"operation_status"`
	CommercialYear  string  `json:"commercial_year"`
}

type Dividend struct {
	Name                         string    `json:"name"`
	Type                         string    `json:"type"`
	Year                         string    `json:"year"`
	TotalStockBonus              float64   `json:"total_stock_bonus"`
	CashDividendPerShareCurrency string    `json:"cash_dividend_per_share_currency"`
	CashDividendPerShare         float64   `json:"cash_dividend_per_share"`
	CumDate                      time.Time `json:"cum_date"`
	ExDate                       time.Time `json:"ex_date"`
	RecordDate                   time.Time `json:"record_date"`
	PaymentDate                  time.Time `json:"payment_date"`
	Ratio1                       int       `json:"ratio_1"`
	Ratio2                       int       `json:"ratio_2"`
	CashDividendCurrency         string    `json:"cash_dividend_currency"`
	CashDividendTotal            float64   `json:"cash_dividend_total"`
}
