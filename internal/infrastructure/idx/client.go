package idx

import (
	"context"
	"encoding/json"
	"fmt"
	"go-stock/internal/shared/rest"
	"net/http"
	"strings"
	"time"
)

type IdxClient interface {
	GetStockList(ctx context.Context) (*StockListResponse, error)
	GetStockSummaryList(ctx context.Context, date string) (*StockSummaryListResponse, error)
	GetBrokerList(ctx context.Context) (*BrokerListResponse, error)
	GetCompanyProfile(ctx context.Context, code string) (*CompanyProfileResponse, error)
	GetFinancialReports(ctx context.Context, period string, year string) (*FinancialReportResponse, error)
}

type idxClient struct {
	restClient           rest.RestClient
	stockListPath        string
	stockSummaryListPath string
	brokerListPath       string
	companyProfilePath   string
	financialReportPath  string
}

type Config struct {
	BaseURL string
	Delay   time.Duration
	Path    Path
}

type Path struct {
	StockList        string
	StockSummaryList string
	BrokerList       string
	CompanyProfile   string
	FinancialReport  string
}

func NewIdxClient(cfg Config, client *http.Client) IdxClient {
	restClient := rest.NewRestClientBuilder().
		WithBaseURL(cfg.BaseURL).
		WithHeader("Alt-Used", "idx.co.id").
		WithHeader("Host", "www.idx.co.id").
		WithHeader("Content-Type", "application/json").
		WithHeader("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.0.0 Safari/537.36").
		WithDelay(cfg.Delay).
		WithHTTPClient(client).
		Build()

	return &idxClient{
		restClient:           restClient,
		stockListPath:        cfg.Path.StockList,
		stockSummaryListPath: cfg.Path.StockSummaryList,
		brokerListPath:       cfg.Path.BrokerList,
		companyProfilePath:   cfg.Path.CompanyProfile,
		financialReportPath:  cfg.Path.FinancialReport,
	}
}

func (c *idxClient) GetStockList(ctx context.Context) (*StockListResponse, error) {
	respBody, statusCode, err := c.restClient.SendRequest(ctx, "GET", c.stockListPath, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("error calling stock list endpoint: %w", err)
	}
	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", statusCode)
	}

	var result StockListResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to decode stock list response: %w", err)
	}
	return &result, nil
}

func (c *idxClient) GetStockSummaryList(ctx context.Context, date string) (*StockSummaryListResponse, error) {
	path := strings.ReplaceAll(c.stockSummaryListPath, "{DATE}", date)
	respBody, statusCode, err := c.restClient.SendRequest(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("error calling stock summary list endpoint: %w", err)
	}
	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", statusCode)
	}

	var result StockSummaryListResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to decode stock summary list response: %w", err)
	}
	return &result, nil
}

func (c *idxClient) GetBrokerList(ctx context.Context) (*BrokerListResponse, error) {
	respBody, statusCode, err := c.restClient.SendRequest(ctx, "GET", c.brokerListPath, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("error calling stock list endpoint: %w", err)
	}
	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", statusCode)
	}

	var result BrokerListResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to decode stock list response: %w", err)
	}
	return &result, nil
}

func (c *idxClient) GetCompanyProfile(ctx context.Context, code string) (*CompanyProfileResponse, error) {
	path := strings.ReplaceAll(c.companyProfilePath, "{CODE}", code)
	respBody, statusCode, err := c.restClient.SendRequest(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("error calling company profile endpoint: %w", err)
	}
	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", statusCode)
	}

	var result CompanyProfileResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to decode company profile response: %w", err)
	}
	return &result, nil
}

func (c *idxClient) GetFinancialReports(ctx context.Context, period string, year string) (*FinancialReportResponse, error) {
	path := strings.ReplaceAll(c.financialReportPath, "{PERIOD}", period)
	path = strings.ReplaceAll(path, "{YEAR}", year)
	respBody, statusCode, err := c.restClient.SendRequest(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("error calling financial report endpoint: %w", err)
	}
	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", statusCode)
	}

	var result FinancialReportResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to decode financial report response: %w", err)
	}
	return &result, nil
}
