package lib_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/waqi-dev-community/go-waqi/lib"
)

const (
	waqiApiKey = "test-API-key"
)

var cityFeedResponse = `
{
	"status": "ok",
	"data": {
			"idx": 7397,
			"aqi": 71,
			"time" : {
					"v":1481396400,
					"s":"2016-12-10 19:00:00",
					"tz":"-06:00"
			},
			"city":{
					"name":"Chi_sp, Illinois",
					"url":"https://aqicn.org/city/usa/illinois/chi_sp/",
					"geo":["41.913600","-87.723900"]
			},
			"iaqi":{
					"pm25":{
							"v":71
					}
			},
			"forecast": {
					"daily": {
							"pm25": [{
									"avg": 154,
									"day": "2020-06-13",
									"max": 157,
									"min": 131
							}]
					}
			}
	}
}
`

func TestFeedService_GetCityFeed_Success(t *testing.T) {
	var (
		expectedResponse lib.CityFeedResponse
		city             = "shanghai"
	)

	FeedService := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		t.Run("URL and request method is as expected", func(t *testing.T) {
			expectedURL := fmt.Sprintf("/feed/%s/?token=%s", city, waqiApiKey)
			assert.Equal(t, http.MethodGet, req.Method)
			assert.Equal(t, expectedURL, req.URL.String())
		})

		var resp lib.CityFeedResponse
		err := payloadStringToStruct(cityFeedResponse, &resp)
		require.Nil(t, err)

		w.WriteHeader(http.StatusOK)
		bb, _ := json.Marshal(resp)
		w.Write(bb)
	}))

	defer FeedService.Close()

	ctx := context.Background()

	c := lib.NewAPIClient(waqiApiKey, &lib.Config{BackendURL: lib.MustParseURL(FeedService.URL)})
	resp, err := c.FeedApi.GetCityFeed(ctx, city)
	require.Nil(t, err)
	assert.NotNil(t, resp)

	t.Run("Response is as expected", func(t *testing.T) {
		err := payloadStringToStruct(cityFeedResponse, &expectedResponse)
		require.Nil(t, err)
		assert.Equal(t, expectedResponse, resp)
	})
}
