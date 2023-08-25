package client

import (
	"io"
	"net/http"
)

func NewNullifyClient(nullifyHost string, apiKey string) *NullifyClient {
	return &NullifyClient{
		httpClient: NewHTTPClient(nullifyHost, apiKey),
	}
}

type NullifyClient struct {
	httpClient *http.Client
}

func (c *NullifyClient) GetNullifyData() ([]byte, error) {
	resp, err := c.httpClient.Get("https://nullify.dev/api/data")
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
