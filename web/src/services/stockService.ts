import { api } from "../api/fetcher";
import type { 
    StockResponse, 
    StockSummaryResponse, 
    FinancialReportResponse, 
    BrokerSummaryResponse,
    StockListApiResponse
} from "../types/stock";

export const stockService = {
    // Get all stocks with pagination
    getAllStocks: async (page: number = 1, limit: number = 20) => {
        const response = await api.get<StockListApiResponse>(`/api/v1/stocks?page=${page}&limit=${limit}`);
        return response;
    },

    // Search stocks by code or name
    searchStocks: async (query: string) => {
        const response = await api.get<StockListApiResponse>(`/api/v1/stocks/search?q=${query}`);
        return response;
    },

    // Get stock by code
    getStockByCode: async (stockCode: string) => {
        const response = await api.get<StockResponse>(`/api/v1/stock?stock_code=${stockCode}`);
        return response;
    },

    // Get stock summaries
    getStockSummaries: async (stockCode: string, startDate: string, endDate: string) => {
        const response = await api.get<StockSummaryResponse[]>(
            `/api/v1/stock/summaries?stock_code=${stockCode}&start_date=${startDate}&end_date=${endDate}`
        );
        return response;
    },

    // Get financial reports
    getFinancialReport: async (stockCode: string, reportPeriod: string, reportYear: string) => {
        const response = await api.get<FinancialReportResponse>(
            `/api/v1/financial_report?stock_code=${stockCode}&report_period=${reportPeriod}&report_year=${reportYear}`
        );
        return response;
    },

    // Get broker summaries
    getBrokerSummaries: async (
        stockCode: string, 
        startDate: string, 
        endDate: string, 
        investorType: string = "ALL", 
        transactionType: string = "ALL"
    ) => {
        const response = await api.get<BrokerSummaryResponse>(
            `/api/v1/brokers/summaries?stock_code=${stockCode}&start_date=${startDate}&end_date=${endDate}&investor_type=${investorType}&transaction_type=${transactionType}`
        );
        return response;
    }
};
