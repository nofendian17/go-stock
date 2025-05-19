package model

type FinancialReportRequest struct {
	StockCode    string `json:"stock_code" validate:"required,len=4"`
	ReportPeriod string `json:"report_period" validate:"required,oneof=TW1 TW2 TW3 Audit"`
	ReportYear   string `json:"report_year" validate:"required,len=4,numeric"`
}

type FinancialReportResponse struct {
	StockCode    string       `json:"stock_code"`
	FileModified string       `json:"file_modified"`
	ReportPeriod string       `json:"report_period"`
	ReportYear   string       `json:"report_year"`
	StockName    string       `json:"stock_name"`
	Attachment   []Attachment `json:"attachment"`
}

type Attachment struct {
	StockCode    string `json:"stock_code"`
	StockName    string `json:"stock_name"`
	FileID       string `json:"file_id"`
	FileModified string `json:"file_modified"`
	FileName     string `json:"file_name"`
	FilePath     string `json:"file_path"`
	FileSize     int    `json:"file_size"`
	FileType     string `json:"file_type"`
	ReportPeriod string `json:"report_period"`
	ReportType   string `json:"report_type"`
	ReportYear   string `json:"report_year"`
}
