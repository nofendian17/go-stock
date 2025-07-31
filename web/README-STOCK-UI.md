# Stock UI Implementation

This is a comprehensive stock trading UI built with React, TypeScript, and Tailwind CSS, integrated with the Go Stock API based on the provided Swagger documentation.

## Features

### 📊 Stock List
- Display all available stocks in a card-based grid layout
- Click on any stock card to view detailed information
- Loading states and error handling

### 📈 Stock Detail Pages
- **Overview**: Company information, major shareholders, business overview
- **Chart**: Interactive price and volume charts using Chart.js
- **Reports**: Financial reports with downloadable attachments
- **Dividends**: Dividend history and payment schedules
- **Brokers**: Broker trading activity with top buyers/sellers analysis

### 🎯 API Integration
- Full integration with all endpoints from the Swagger documentation
- Type-safe API calls with proper error handling
- Date range selection for historical data

## Components

### Core Components
- `StockCard.tsx` - Reusable stock display card
- `StockChart.tsx` - Interactive price/volume charts
- `FinancialReportCard.tsx` - Financial report display
- `BrokerSummaryTable.tsx` - Broker activity analysis

### Pages
- `StockList.tsx` - Main stock listing page
- `StockDetail.tsx` - Detailed stock information with tabs

### Services
- `stockService.ts` - API service layer for all stock operations
- `fetcher.ts` - HTTP client configuration

## API Endpoints Used

1. **GET /api/v1/stocks** - List all stocks
2. **GET /api/v1/stock?stock_code={code}** - Get stock details
3. **GET /api/v1/stock/summaries** - Get stock price summaries
4. **GET /api/v1/financial_report** - Get financial reports
5. **GET /api/v1/brokers/summaries** - Get broker trading activity

## Usage

### Running the Application

1. Install dependencies:
```bash
cd web
npm install
```

2. Start the development server:
```bash
npm run dev
```

3. Ensure the Go Stock API is running on `http://localhost:3000`

### Environment Configuration

Create a `.env` file in the web directory:
```
VITE_API_BASE_URL=http://localhost:3000
```

## Navigation

- `/` - Home page
- `/stocks` - Stock list
- `/stocks/:stockCode` - Stock detail page

## Features Overview

### Stock List Page
- Responsive grid layout
- Company logos and basic information
- Click navigation to detail pages

### Stock Detail Page
- **Overview Tab**: Company profile, shareholders, subsidiaries
- **Chart Tab**: Interactive price charts with date range selection
- **Reports Tab**: Financial reports with downloadable attachments
- **Dividends Tab**: Dividend history and payment schedules
- **Brokers Tab**: Trading activity by brokers with top buyers/sellers

## Technology Stack

- **Frontend**: React 18 with TypeScript
- **Styling**: Tailwind CSS
- **Charts**: Chart.js with React Chart.js 2
- **Routing**: React Router
- **HTTP Client**: Fetch API with custom wrapper
- **Build Tool**: Vite

## Type Safety

All API responses are fully typed with TypeScript interfaces based on the Swagger schema:
- `StockResponse`
- `StockSummaryResponse`
- `FinancialReportResponse`
- `BrokerSummaryResponse`
- And supporting interfaces for all nested data structures
