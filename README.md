# go-stock

A stock data management system built with Go, providing RESTful APIs for stock information, brokers, financial reports, and more. Uses clean architecture with MongoDB for persistence and integrates external financial data sources.

> **Note:** This project is for educational purposes only.

## Features
- RESTful API endpoints for stock data, brokers, and financial reports
- MongoDB persistence layer with repository pattern
- Clean architecture separation of concerns
- External data integration via IDX and Indopremier 
- Health check endpoint
- Cron job support for scheduled data synchronization
- Structured error handling and JSON responses

## Tech Stack
- Go 1.20+
- MongoDB
- RESTful API 
- Clean Architecture
- Docker support

## Setup

### Prerequisites
- Docker
- MongoDB instance (local or cloud)
- Go 1.20+ installed

### Quick Start (Docker)
```bash
# Build and run with Docker Compose
docker-compose up -d --build
```

### Manual Setup
1. Clone repository
2. Copy `example.config.yaml` to `config.yaml`
3. Update MongoDB connection settings
4. Install dependencies
```bash
go mod download
```
5. Run application
```bash
go run main.go
```

## API Endpoints
Base URL: `http://localhost:8080/`

For a complete API specification, please see the [Swagger documentation](http://localhost:8080/swagger/index.html).

## Scheduled Tasks
- **Stock Data Synchronization**: Runs from config in `cron_job` to refresh stock data from IDX API.

```
internal/
├── config/      - Configuration management
├── delivery/    - HTTP and cron endpoints
├── infrastructure/ - Database/API clients
├── repository/  - Data access layer
├── usecase/     - Business logic
└── entity/      - Domain models
```

## Contribution
1. Fork the repository
2. Create feature branch (`git checkout -b feature/something`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin feature/something`)
5. Create new Pull Request

## License
MIT License © 2025
