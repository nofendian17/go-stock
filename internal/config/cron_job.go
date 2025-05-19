package config

type CronJob struct {
	UpdateStockList        string `mapstructure:"update_stock_list"`
	UpdateStockSummaryList string `mapstructure:"update_stock_summary_list"`
	UpdateBrokerList       string `mapstructure:"update_broker_list"`
	UpdateFinancialReport  string `mapstructure:"update_financial_report"`
}
