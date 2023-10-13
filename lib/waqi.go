package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
)

const (
	WaqiURL = "https://api.waqi.info"
)

type Config struct {
	BackendURL *url.URL
	HttpClient *http.Client
}

type APIClient struct {
	apiKey string
	config *Config
	common service

	// Api Service
	FeedApi *FeedService
}

type service struct {
	client *APIClient
}

func NewAPIClient(apiKey string, cfg *Config) *APIClient {
	cfg.BackendURL = buildBackendURL(cfg)

	if cfg.HttpClient == nil {
		cfg.HttpClient = &http.Client{Timeout: 20 * time.Second}
	}

	c := &APIClient{apiKey: apiKey}
	c.config = cfg
	c.common.client = c

	// API Services
	c.FeedApi = (*FeedService)(&c.common)
	return c
}

func (c APIClient) sendRequest(req *http.Request, resp interface{}) (*http.Response, error) {
	req.Header.Set("Content-Type", "application/json")
	tokenQuery := req.URL.Query()
	tokenQuery.Add("token", c.apiKey)

	req.URL.RawQuery = tokenQuery.Encode()

	res, err := c.config.HttpClient.Do(req)
	if err != nil {
		return res, errors.Wrap(err, "failed to execute request")
	}

	body, _ := io.ReadAll(res.Body)
	defer res.Body.Close()

	if res.StatusCode >= http.StatusMultipleChoices {
		return res, errors.Errorf(
			`request was not successful, status code %d, %s`, res.StatusCode,
			string(body),
		)
	}

	if string(body) == "" {
		resp = map[string]string{}
		return res, nil
	}

	err = c.decode(&resp, body)
	if err != nil {
		return res, errors.Wrap(err, "unable to unmarshal response body")
	}

	return res, nil
}

func (c APIClient) decode(v interface{}, b []byte) (err error) {
	if err = json.Unmarshal(b, v); err != nil {
		return err
	}
	return nil
}

func buildBackendURL(cfg *Config) *url.URL {
	if cfg.BackendURL == nil {
		rawURL := fmt.Sprintf("%s", WaqiURL)
		return MustParseURL(rawURL)
	}

	return cfg.BackendURL
}

func MustParseURL(rawURL string) *url.URL {
	u, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}

	return u
}
