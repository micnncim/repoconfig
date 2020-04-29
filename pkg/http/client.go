package http

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"go.uber.org/zap"
)

// Client is a HTTP client.
type Client struct {
	httpClient *http.Client

	endpoint           *url.URL
	username, password string

	logger *zap.Logger
}

// Option represents a option for Client constructor.
type Option func(*Client)

// WithTimeout returns Option which sets an optional timeout for Client.
func WithTimeout(t time.Duration) Option {
	return func(c *Client) { c.httpClient.Timeout = t }
}

// WithLogger returns Option which sets an optional logger for Client.
func WithLogger(l *zap.Logger) Option {
	return func(c *Client) { c.logger = l.Named("http") }
}

// NewClient returns a Client with initializations.
func NewClient(endpoint string, opts ...Option) (*Client, error) {
	if endpoint == "" {
		return nil, errors.New("endpoint for client must be set")
	}
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	httpClient := &http.Client{
		Timeout: time.Minute,
	}
	c := &Client{
		httpClient: httpClient,
		endpoint:   u,
		logger:     zap.NewNop(),
	}

	for _, opt := range opts {
		opt(c)
	}

	return c, nil
}

// DoRequest does HTTP request, wrapping http.Client.Do. The body is assumed to be JSON.
func (c *Client) DoRequest(ctx context.Context, method, path string, headers map[string]string, body interface{}) (*http.Response, error) {
	logger := c.logger.With(zap.String("method", method), zap.String("path", path))

	var buf io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		logger.Debug("create http request body", zap.Any("payload", string(b)))
		buf = bytes.NewBuffer(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, fmt.Sprintf("%s/%s", c.endpoint.String(), path), buf)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	return c.httpClient.Do(req)
}
