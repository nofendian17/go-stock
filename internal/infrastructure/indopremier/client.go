package indopremier

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"go-stock/internal/shared/rest"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type IndopremierClient interface {
	GetBrokerSummary(ctx context.Context, stockCode, startDate, endDate, investorType, board string) (*GetBrokerSummaryResponse, error)
}

type indopremierClient struct {
	restClient        rest.RestClient
	BrokerSummaryPath string
}

type Config struct {
	BaseURL string
	Delay   time.Duration
	Path    Path
}

type Path struct {
	BrokerSummary string
}

// NewIndopremierClient initializes a new IndopremierClient.
func NewIndopremierClient(cfg Config, client *http.Client) IndopremierClient {
	restClient := rest.NewRestClientBuilder().
		WithBaseURL(cfg.BaseURL).
		WithHeader("Content-Type", "text/html").
		WithHeader("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.0.0 Safari/537.36").
		WithDelay(cfg.Delay).
		WithHTTPClient(client).
		Build()

	return &indopremierClient{
		restClient:        restClient,
		BrokerSummaryPath: cfg.Path.BrokerSummary,
	}
}

// GetBrokerSummary fetches broker summary data.
func (c *indopremierClient) GetBrokerSummary(ctx context.Context, stockCode, startDate, endDate, investorType, board string) (*GetBrokerSummaryResponse, error) {
	sd, err := c.parseDate(startDate)
	if err != nil {
		return nil, fmt.Errorf("error parsing start date: %w", err)
	}

	ed, err := c.parseDate(endDate)
	if err != nil {
		return nil, fmt.Errorf("error parsing end date: %w", err)
	}

	// Replace placeholders in the path template
	path := strings.NewReplacer(
		"{CODE}", stockCode,
		"{START_DATE}", startDate,
		"{END_DATE}", endDate,
		"{INVESTOR_TYPE}", investorType,
		"{BOARD}", board,
	).Replace(c.BrokerSummaryPath)

	// Send the HTTP request
	respBody, statusCode, err := c.restClient.SendRequest(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("error calling broker summary endpoint: %w", err)
	}
	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", statusCode)
	}

	// Parse the response
	result, err := c.brokerSummaryParse(respBody)
	if err != nil {
		return nil, err
	}

	result.StockCode = stockCode
	result.StartDate = sd
	result.EndDate = ed

	return result, nil
}

// brokerSummaryParse parses the HTML response into a structured format.
func (c *indopremierClient) brokerSummaryParse(html []byte) (*GetBrokerSummaryResponse, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(html)))
	if err != nil {
		return &GetBrokerSummaryResponse{}, fmt.Errorf("failed to parse HTML: %w", err)
	}

	var result GetBrokerSummaryResponse

	// Flag untuk cek apakah ada baris valid
	var hasValidRow bool

	doc.Find("table.table-summary tbody tr").Each(func(i int, s *goquery.Selection) {
		tds := s.Find("td")
		if tds.Length() < 9 {
			return // Skip rows with insufficient data
		}

		// Ambil isi teks
		buyerCode := strings.TrimSpace(tds.Eq(0).Text())
		sellerCode := strings.TrimSpace(tds.Eq(5).Text())

		// Cek jika tidak ada broker code di buyer & seller
		if buyerCode == "" && sellerCode == "" {
			return // kosong, lewati
		}

		hasValidRow = true // ada setidaknya satu baris valid

		buyer := BrokerSummaryData{
			BrokerCode: buyerCode,
			Lot:        c.parseNumber(tds.Eq(1).Text()),
			Val:        strings.TrimSpace(tds.Eq(2).Text()),
			Avg:        c.parseNumber(tds.Eq(3).Text()),
		}

		seller := BrokerSummaryData{
			BrokerCode: sellerCode,
			Lot:        c.parseNumber(tds.Eq(6).Text()),
			Val:        strings.TrimSpace(tds.Eq(7).Text()),
			Avg:        c.parseNumber(tds.Eq(8).Text()),
		}

		result.Buyers = append(result.Buyers, buyer)
		result.Sellers = append(result.Sellers, seller)
	})

	// Parsing footer summary
	doc.Find("table.table-summary tfoot tr th div span").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())

		switch {
		case strings.HasPrefix(text, "T. Val"):
			result.Summary.TotalVal = c.extractValue(text)
		case strings.HasPrefix(text, "F. NVal"):
			result.Summary.ForeignNetVal = c.extractValue(text)
		case strings.HasPrefix(text, "T.Lot"):
			result.Summary.TotalLot = c.parseNumber(c.extractValue(text))
		case strings.HasPrefix(text, "Avg"):
			result.Summary.Avg = c.parseNumber(c.extractValue(text))
		}
	})

	// Jika tidak ada baris valid dan summary juga kosong, anggap datanya kosong
	if !hasValidRow && result.Summary.TotalVal == "0" && result.Summary.TotalLot == 0 && result.Summary.Avg == 0 {
		return nil, fmt.Errorf("empty broker summary data")
	}

	return &result, nil
}

// parseNumber converts a string representation of a number into a float64.
func (c *indopremierClient) parseNumber(val string) float64 {
	fields := strings.Fields(strings.ReplaceAll(strings.TrimSpace(val), ",", ""))
	if len(fields) == 0 {
		return 0
	}
	num, _ := strconv.ParseFloat(fields[0], 64) // Ignore errors, return 0 if invalid
	return num
}

// extractValue ...
func (c *indopremierClient) extractValue(text string) string {
	parts := strings.SplitN(text, ":", 2)
	if len(parts) < 2 {
		return ""
	}
	return strings.TrimSpace(parts[1])
}

// parseDate ...
func (c *indopremierClient) parseDate(date string) (time.Time, error) {
	t, err := time.Parse("01/02/2006", date)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}
