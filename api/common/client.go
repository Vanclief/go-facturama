package common

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/vanclief/ez"
)

const (
	// ProductionBaseURL is the base URL for the production API
	ProductionBaseURL = "https://api.facturama.mx"
	// SandboxBaseURL is the base URL for the sandbox API
	SandboxBaseURL = "https://apisandbox.facturama.mx"

	// DefaultTimeout is the default timeout for HTTP requests
	DefaultTimeout = 30 * time.Second
)

// Environment represents the API environment
type Environment string

const (
	// Production environment
	Production Environment = "production"
	// Sandbox environment
	Sandbox Environment = "sandbox"
)

// Client represents a Facturama API client
type Client struct {
	// HTTPClient is the underlying HTTP client
	HTTPClient *http.Client

	// BaseURL is the base URL for API requests
	BaseURL string

	// Authentication credentials
	Username string
	Password string
}

// Option is a function that configures a Client
type Option func(*Client)

// WithTimeout sets the HTTP client timeout
func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.HTTPClient.Timeout = timeout
	}
}

// WithHTTPClient sets a custom HTTP client
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.HTTPClient = httpClient
	}
}

// WithEnvironment sets the API environment
func WithEnvironment(env Environment) Option {
	return func(c *Client) {
		switch env {
		case Production:
			c.BaseURL = ProductionBaseURL
		case Sandbox:
			c.BaseURL = SandboxBaseURL
		}
	}
}

// WithBaseURL sets a custom base URL
func WithBaseURL(baseURL string) Option {
	return func(c *Client) {
		c.BaseURL = baseURL
	}
}

// NewClient creates a new Facturama API client
func NewClient(username, password string, options ...Option) *Client {
	client := &Client{
		HTTPClient: &http.Client{
			Timeout: DefaultTimeout,
		},
		BaseURL:  SandboxBaseURL, // Default to sandbox
		Username: username,
		Password: password,
	}

	// Apply options
	for _, option := range options {
		option(client)
	}

	return client
}

// Request makes an HTTP request to the Facturama API
func (c *Client) Request(ctx context.Context, method, path string, body interface{}, response interface{}) error {
	const op = "common.Request"

	// Create full URL
	url := c.BaseURL + path

	// Create request body if provided
	var bodyReader io.Reader
	if body != nil {
		bodyData, err := json.Marshal(body)
		if err != nil {
			return ez.New(op, ez.EINTERNAL, "Error marshaling request body request", err)
		}
		bodyReader = bytes.NewReader(bodyData)
	}

	// Create request
	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return ez.New(op, ez.EINTERNAL, "Error creating request", err)
	}

	// Create Basic Authentication header
	req.SetBasicAuth(c.Username, c.Password)

	// Add other headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Execute request
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return ez.New(op, ez.EINTERNAL, "Error executing request", err)
	}
	defer resp.Body.Close()

	// Read response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return ez.New(op, ez.EINTERNAL, "Error reading response body", err)
	}

	// Check for error response
	if resp.StatusCode >= 400 {

		var errResp ErrorResponse
		if err := json.Unmarshal(responseBody, &errResp); err != nil {
			// fallback to raw string if JSON is malformed
			return &APIError{
				RawBody:    string(responseBody),
				StatusCode: resp.StatusCode,
			}
		}

		return &APIError{
			ErrorResponse: errResp,
			RawBody:       string(responseBody),
			StatusCode:    resp.StatusCode,
		}

	}

	// Parse response if provided
	if response != nil && len(responseBody) > 0 {
		if err := json.Unmarshal(responseBody, response); err != nil {
			return ez.New(op, ez.EINTERNAL, "Error unmarshaling response", err)
		}
	}

	return nil
}

// Get makes a GET request to the API
func (c *Client) Get(ctx context.Context, path string, response interface{}) error {
	return c.Request(ctx, http.MethodGet, path, nil, response)
}

// Post makes a POST request to the API
func (c *Client) Post(ctx context.Context, path string, body, response interface{}) error {
	return c.Request(ctx, http.MethodPost, path, body, response)
}

// Put makes a PUT request to the API
func (c *Client) Put(ctx context.Context, path string, body, response interface{}) error {
	return c.Request(ctx, http.MethodPut, path, body, response)
}

// Delete makes a DELETE request to the API
func (c *Client) Delete(ctx context.Context, path string) error {
	return c.Request(ctx, http.MethodDelete, path, nil, nil)
}

// APIError represents an error returned by the Facturama API
type APIError struct {
	ErrorResponse // the structured error payload

	// Original error response body
	RawBody string

	// HTTP status code
	StatusCode int
}

// ErrorResponse models the JSON payload returned on HTTP 4xx/5xx
type ErrorResponse struct {
	Message    string     `json:"Message"`
	ModelState ModelState `json:"ModelState"`
}

// ModelState captures all of the possible ModelState arrays.
// Add or omit fields as needed; use `omitempty` so zero slices disappear.
type ModelState struct {
	Message     []string `json:"Message,omitempty"`
	Certificate []string `json:"Certificate,omitempty"`
	Key         []string `json:"Key,omitempty"`
}

// Error makes APIError implement the error interface.
func (e *APIError) Error() string {
	// you can format this however you like—here’s a concise version:
	return fmt.Sprintf(
		"Facturama API error %d: %s — %v",
		e.StatusCode,
		e.Message,
		e.ModelState,
	)
}
