package client

import "net/http"

type authTransport struct {
	token     string
	transport http.RoundTripper
}

func (t *authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer "+t.token)
	return t.transport.RoundTrip(req)
}

func NewHTTPClient(nullifyHost string, apiKey string) *http.Client {
	return &http.Client{
		Transport: &authTransport{
			token:     apiKey,
			transport: http.DefaultTransport,
		},
	}
}
