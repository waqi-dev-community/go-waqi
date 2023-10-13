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

**For City Feed:**

```golang
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

	city := "shanghai"
	cityFeed, err := waqiClient.FeedApi.GetCityFeed(ctx, city)

	if err != nil {
		log.Fatal("city feed error: ", err.Error())
		return
	}

	fmt.Println(cityFeed)
}
```