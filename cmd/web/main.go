package main

import (
	"fmt"

	"github.com/eklipsan/forumeather/pkg/meteoblue"
)


func main() {
	// test := meteoblue.NewLocation("60.1695", "24.9354", "TestLocation")
	makeevka := meteoblue.NewLocation(48.0478,37.9258, "Makeevka")

	result, err := makeevka.GetCurrentForecast()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)

}