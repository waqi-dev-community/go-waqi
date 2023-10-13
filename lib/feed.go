package lib

import (
	"context"
	"net/http"
)

type FeedService service

func (c *FeedService) GetCityFeed(ctx context.Context, city string) (CityFeedResponse, error) {
	var resp CityFeedResponse
	URL := c.client.config.BackendURL.JoinPath("feed", city, "/")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL.String(), http.NoBody)
	if err != nil {
		return resp, err
	}

	_, err = c.client.sendRequest(req, &resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
