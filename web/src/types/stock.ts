export interface StockResponse {
  code: string;
  name: string;
  share: number;
  listing_date: string;
  board: string;
  profiles: Profile[];
  secretaries: Secretary[];
  directors: Director[];
  commissioners: Commissioner[];
  audit_committees: AuditCommittee[];
  shareholders: Shareholder[];
  subsidiaries: Subsidiary[];
  dividends: Dividend[];
}

export interface Profile {
  address: string;
  bae: string;
  industry: string;
  sub_industry: string;
  email: string;
  fax: string;
  main_business: string;
  stock_code: string;
  stock_name: string;
  tin: string;
  sector: string;
  sub_sector: string;
  listing_date: string;
  phone: string;
  website: string;
  status: number;
  logo: string;
}

export interface Secretary {
  name: string;
  phone_number: string;
  website: string;
  email: string;
  fax: string;
  mobile_number: string;
}

export interface Director {
  name: string;
  position: string;
  is_affiliated: boolean;
}

export interface Commissioner {
  name: string;
  position: string;
  is_independent: boolean;
}

export interface AuditCommittee {
  name: string;
  position: string;
}

export interface Shareholder {
  share: number;
  category: string;
  name: string;
  is_controller: boolean;
  percentage: number;
}

export interface Subsidiary {
  business_fields: string;
  total_asset: number;
  location: string;
  currency: string;
  name: string;
  percentage: number;
  units: string;
  operation_status: string;
  commercial_year: string;
}

export interface Dividend {
  name: string;
  type: string;
  year: string;
  total_stock_bonus: number;
  cash_dividend_per_share_currency: string;
  cash_dividend_per_share: number;
  cum_date: string;
  ex_date: string;
  record_date: string;
  payment_date: string;
  ratio_1: number;
  ratio_2: number;
  cash_dividend_currency: string;
  cash_dividend_total: number;
}

export interface StockSummaryResponse {
  id_stock_summary: number;
  stock_code: string;
  stock_name: string;
  date: string;
  previous: number;
  open_price: number;
  high: number;
  low: number;
  close: number;
  change: number;
  percentage: number;
  volume: number;
  value: number;
  frequency: number;
  index_individual: number;
  offer: number;
  offer_volume: number;
  bid: number;
  bid_volume: number;
  listed_shares: number;
  tradeble_shares: number;
  weight_for_index: number;
  delisting_date: string;
  first_trade: number;
  foreign_buy: number;
  foreign_sell: number;
  non_regular_volume: number;
  non_regular_value: number;
  non_regular_frequency: number;
  persen: number;
  remarks: string;
}

export interface FinancialReportResponse {
  stock_code: string;
  stock_name: string;
  report_year: string;
  report_period: string;
  file_modified: string;
  attachment: Attachment[];
}

export interface Attachment {
  file_id: string;
  file_name: string;
  file_path: string;
  file_size: number;
  file_type: string;
  file_modified: string;
  report_year: string;
  report_period: string;
  report_type: string;
  stock_code: string;
  stock_name: string;
}

export interface BrokerSummaryResponse {
  stock_code: string;
  start_date: string;
  end_date: string;
  buyers: BrokerSummaryData[];
  sellers: BrokerSummaryData[];
  summary: Summary;
}

export interface BrokerSummaryData {
  broker_code: string;
  lot: number;
  val: string;
  avg: number;
}

export interface Summary {
  total_lot: number;
  total_val: string;
  foreign_net_val: string;
  avg: number;
}

// Standard API response wrapper
export interface ApiResponse<T> {
  code: number;
  message: string;
  data: T;
}

// Pagination response interface
export interface PaginationResponse {
  data: StockResponse[];
  total: number;
  page: number;
  limit: number;
  total_pages: number;
}

// Combined type for the actual API response
export type StockListApiResponse = ApiResponse<PaginationResponse>;
