package lib

const (
	HTTPStatusOk       = 200
	HTTPStatusCreated  = 201
	HTTPStatusNoChange = 204
	HTTPRedirectOk     = 300
)

type TokenQuery struct {
	Token string `json:"token"`
}

type QueryBuilder interface {
	BuildQuery() string
}

type GeoFeedResponse struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}

type CityFeedResponseData struct {
	IDX  int `json:"idx"`
	AQI  int `json:"aqi"`
	Time struct {
		v  int64
		S  string
		TZ string
	} `json:"time"`
	City struct {
		Name string
		Geo  []string
		URL  string
	} `json:"city"`
	IAQI     interface{} `json:"iaqi"`
	Forecast interface{} `json:"forecast"`
}

type CityFeedResponse struct {
	Status string               `json:"status"`
	Data   CityFeedResponseData `json:"data"`
}
