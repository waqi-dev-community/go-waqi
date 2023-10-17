# Go WAQI Client

Go client library for the World Air Quality Index (WAQI) APIs. See documentation [here](https://aqicn.org/json-api/doc/).
API modules supported - City feed.

### Installation

You can include this package in your Go project using npm or yarn:

```shell
go get github.com/waqi-dev-community/go-waqi
```

### Get API key

Sign up for an API key [here](https://aqicn.org/data-platform/token/)

### Making Requests

Making Requests
```golang

import (
	waqi "github.com/waqi-dev-community/go-waqi/lib"
)

apiKey := "<REPLACE_WITH_YOUR_API_KEY>"
waqiClient := waqi.NewAPIClient(apiKey, &waqi.Config{})
```

**For City Feed:**

```golang


city := "shanghai"
cityFeed, err := waqiClient.FeedApi.GetCityFeed(ctx, city)

if err != nil {
	log.Fatal("city feed error: ", err.Error())
	return
}

fmt.Println(cityFeed)
```

**For Lat/Lng based Geolocalized Feed:**

```golang

lat := 32.455
lng := 10.322
geoFeed, err := waqiClient.FeedApi.GeoFeed(ctx, lng, lat)

if err != nil {
	log.Fatal("geo feed error: ", err.Error())
	return
}

fmt.Println(geoFeed)
```

**For IP based Geolocalized Feed:**

```golang
geoFeed, err := waqiClient.FeedApi.GeoFeed(ctx)

if err != nil {
	log.Fatal("ip feed error: ", err.Error())
	return
}

fmt.Println(geoFeed)
```
