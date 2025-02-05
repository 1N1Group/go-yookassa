package v3

import (
	"encoding/json"
	"fmt"
	"time"

	"resty.dev/v3"
)

const (
	baseUrl = "https://api.yookassa.ru/v3"
)

type Client struct {
	client    *resty.Client
	shopID    string
	secretKey string
	testMode  bool
	retries   int
	timeout   time.Duration
	// API
	Payments PaymentsService
}

type ClientConfig struct {
	TestMode bool
	Retries  int
	Timeout  time.Duration
}

func DefaultConfig() *ClientConfig {
	return &ClientConfig{
		TestMode: false,
		Retries:  0,
		Timeout:  1 * time.Minute,
	}
}

func New(shopID, secretKey string, config *ClientConfig) *Client {
	httpClient := resty.New()
	httpClient.SetBasicAuth(shopID, secretKey)
	httpClient.SetBaseURL(baseUrl)
	httpClient.SetRetryCount(config.Retries)
	httpClient.SetTimeout(config.Timeout)

	if config == nil {
		config = DefaultConfig()
	}

	c := &Client{
		client:    httpClient,
		shopID:    shopID,
		secretKey: secretKey,
		testMode:  config.TestMode,
		retries:   config.Retries,
		timeout:   config.Timeout,
	}

	c.Payments = &PaymentsServiceImpl{client: c}

	return c
}

func (c *Client) request(method string, urlStr string, queryParams interface{}, requestBody interface{}) (interface{}, error) {

	req := c.client.R().
		SetContentType("application/json").
		SetMethod(method).
		SetURL(urlStr)

	if queryParams != nil {
		var query map[string]string

		line, err := json.Marshal(queryParams)
		json.Unmarshal(line, &query)

		if err != nil {
			return nil, err
		}

		req = req.SetQueryParams(query)
	}

	if requestBody != nil {

		var body map[string]interface{}

		line, err := json.Marshal(requestBody)
		json.Unmarshal(line, &body)

		if err != nil {
			return nil, err
		}

		body["test"] = c.testMode

		req = req.SetBody(body)
	}

	resp, err := req.Send()

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("API error: %s", resp.String())
	}

	return resp.Result(), nil
}
