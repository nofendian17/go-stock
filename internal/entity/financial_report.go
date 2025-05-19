package entity

type FinancialReport struct {
	StockCode    string       `bson:"stock_code"`
	FileModified string       `bson:"file_modified"`
	ReportPeriod string       `bson:"report_period"`
	ReportYear   string       `bson:"report_year"`
	StockName    string       `bson:"stock_name"`
	Attachment   []Attachment `bson:"attachment"`
}

type Attachment struct {
	StockCode    string `bson:"stock_code"`
	StockName    string `bson:"stock_name"`
	FileID       string `bson:"file_id"`
	FileModified string `bson:"file_modified"`
	FileName     string `bson:"file_name"`
	FilePath     string `bson:"file_path"`
	FileSize     int    `bson:"file_size"`
	FileType     string `bson:"file_type"`
	ReportPeriod string `bson:"report_period"`
	ReportType   string `bson:"report_type"`
	ReportYear   string `bson:"report_year"`
}
