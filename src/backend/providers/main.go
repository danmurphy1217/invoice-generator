package providers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

type HttpClient interface {
	/*
	interface for sending HTTP requests

	currently, must only implement POST request
	*/

	// Post information to the URL
	Post(url string, contentType string, body []byte, headers map[string]string) ([]byte, error)
}


type httpClient struct {
    client *http.Client
}

func NewHttpClient(timeout time.Duration) HttpClient {
    return &httpClient{
        client: &http.Client{
            Timeout: timeout,
        },
    }
}

func (c *httpClient) Post(url string, contentType string, body []byte, headers map[string]string) ([]byte, error) {
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
    if err != nil {
        return nil, fmt.Errorf("failed to create request: %w", err)
    }

    req.Header.Set("Content-Type", contentType)
    for key, value := range headers {
        req.Header.Set(key, value)
    }

    resp, err := c.client.Do(req)
    if err != nil {
        return nil, fmt.Errorf("failed to execute request: %w", err)
    }
    defer resp.Body.Close()

    responseBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("failed to read response body: %w", err)
    }

    if resp.StatusCode >= 300 {
        return responseBody, fmt.Errorf("received non-success status code: %d", resp.StatusCode)
    }

    return responseBody, nil
}