package main

import (
	"context"
	"fmt"
	"log"

	waqi "github.com/waqi-dev-community/go-waqi/lib"
)

func main() {
	apiKey := "<REPLACE_WITH_YOUR_API_KEY>"

	ctx := context.Background()
	waqiClient := waqi.NewAPIClient(apiKey, &waqi.Config{})

	// City Feed
	city := "shanghai"
	cityFeed, err := waqiClient.FeedApi.GetCityFeed(ctx, city)

	if err != nil {
		log.Fatal("city feed error: ", err.Error())
		return
	}

	fmt.Println(cityFeed)
}
