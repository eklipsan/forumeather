package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/eklipsan/forumeather/pkg/meteoblue"
)


func main() {

	// resp, err := http.Get("https://my.meteoblue.com/packages/current?apikey=yZnQoPluHmAvFVxj&lat=60.1695&lon=24.9354&asl=26&tz=utc&format=json")

	url, err := meteoblue.CreateURL()

	resp, err := http.Get(url)

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(string(result))
	}

	myStruct, _ := meteoblue.UnmarshalCurrentForecast(result)

	fmt.Println(myStruct.Units.Time)
	defer resp.Body.Close()
}