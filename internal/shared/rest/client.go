package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"time"
)

type RestClient interface {
	SendRequest(ctx context.Context, method, url string, body interface{}, headers map[string]string) (responseBody []byte, statusCode int, err error)
}

type restClient struct {
	httpClient *http.Client
	delay      time.Duration
	baseURL    string
	headers    map[string]string
}

type RestClientBuilder struct {
	client *restClient
}

func NewRestClientBuilder() *RestClientBuilder {
	return &RestClientBuilder{
		client: &restClient{
			httpClient: &http.Client{},
			headers:    make(map[string]string),
		},
	}
}

func (b *RestClientBuilder) WithBaseURL(baseURL string) *RestClientBuilder {
	b.client.baseURL = baseURL
	return b
}

func (b *RestClientBuilder) WithHeader(key, value string) *RestClientBuilder {
	b.client.headers[key] = value
	return b
}

func (b *RestClientBuilder) WithHTTPClient(httpClient *http.Client) *RestClientBuilder {
	b.client.httpClient = httpClient
	return b
}

func (b *RestClientBuilder) WithDelay(delay time.Duration) *RestClientBuilder {
	b.client.delay = delay
	return b
}

func (b *RestClientBuilder) Build() RestClient {
	return b.client
}

func (r *restClient) SendRequest(ctx context.Context, method, path string, body interface{}, headers map[string]string) (responseBody []byte, statusCode int, err error) {
	var reqBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, 0, err
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	fullURL := r.baseURL + path
	req, err := http.NewRequestWithContext(ctx, method, fullURL, reqBody)
	if err != nil {
		return nil, 0, err
	}

	for key, value := range r.headers {
		req.Header.Set(key, value)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Log the request
	slog.Info("HTTP Request",
		"method", method,
		"url", fullURL,
		"headers", r.headers,
		"body", body,
	)

	time.Sleep(r.delay)

	resp, err := r.httpClient.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	responseBody, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	// Log the response
	slog.Info("HTTP Response",
		"statusCode", resp.StatusCode,
		"headers", resp.Header,
		"body", string(responseBody),
	)

	return responseBody, resp.StatusCode, nil
}
