package lib

import (
	"context"
	"fmt"
	"net/http"
)

type FeedService service

func (c *FeedService) GetCityFeed(ctx context.Context, city string) (CityFeedResponse, error) {
	var resp CityFeedResponse
	_city := fmt.Sprintf("%s/", city)
	URL := c.client.config.BackendURL.JoinPath("feed", _city)

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

func (c *FeedService) GeoFeed(ctx context.Context, lng float64, lat float64) (GeoFeedResponse, error) {
	var resp GeoFeedResponse
	format := fmt.Sprintf("geo:%v;%v/", lng, lat)
	URL := c.client.config.BackendURL.JoinPath("feed", format)

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

func (c *FeedService) IPFeed(ctx context.Context) (GeoFeedResponse, error) {
	var resp GeoFeedResponse
	URL := c.client.config.BackendURL.JoinPath("feed", "here/")

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
