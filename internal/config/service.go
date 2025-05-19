package config

type Service struct {
	IDXService         `mapstructure:"idx_service"`
	IndoPremierService `mapstructure:"indo_premier_service"`
}

type IDXService struct {
	BaseURL string `mapstructure:"base_url"`
	Delay   int    `mapstructure:"delay"`
	Path    struct {
		StockList        string `mapstructure:"stock_list"`
		StockSummaryList string `mapstructure:"stock_summary_list"`
		BrokerList       string `mapstructure:"broker_list"`
		CompanyProfile   string `mapstructure:"company_profile"`
		FinancialReport  string `mapstructure:"financial_report"`
	}
}

type IndoPremierService struct {
	BaseURL string `mapstructure:"base_url"`
	Path    struct {
		BrokerSummary string `mapstructure:"broker_summary"`
	}
}
