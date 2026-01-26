package meteoblue

import (
	"fmt"
	"io"
	"net/http"

	"github.com/eklipsan/forumeather/pkg/config"
)

type MeteobluePackage string

const (
	CurrentPackage MeteobluePackage = "current"
	Basic15MinutePackage MeteobluePackage = "basic-15min"
	Basic1HourPackage MeteobluePackage = "basic-1h"
	Basic3HourPackage MeteobluePackage = "basic=3h"
	BasicDayPackage MeteobluePackage = "basic-day"
)

type weatherLocation struct {
	Lat float64
	Lon float64
	Name string
	CurrentURL string
}

func NewLocation(lat, lon float64, name string) *weatherLocation {
	return &weatherLocation{
		Lat: lat,
		Lon: lon,
		Name: name,
	}
}


func (w *weatherLocation) GetCurrentForecast() (CurrentForecast, error) {
	currentUrl, err := w.createURL(CurrentPackage)
	if err != nil {
		return CurrentForecast{}, err
	}

	req, err := http.Get(currentUrl)
	if err != nil || req.StatusCode != http.StatusOK {
		return CurrentForecast{}, err
	}
	defer req.Body.Close()

	bytesResult, err := io.ReadAll(req.Body)
	if err != nil {
		return CurrentForecast{}, err
	}

	structResult, err := UnmarshalCurrentForecast(bytesResult)
	if err != nil {
		return CurrentForecast{}, err
	}

	return structResult, nil

}


func (w *weatherLocation) createURL(PackageName MeteobluePackage) (string, error) {
	var meteoblueURL string
	configURL, err := config.LoadMeteoblueConfig()

	if err != nil {
		return "", err
	}
	meteoblueURL = configURL.URL
	meteoblueURL += string(PackageName)
	meteoblueURL += "?apikey=" + configURL.APIKey
	strLat := fmt.Sprintf("%.4f", w.Lat)
	strLon := fmt.Sprintf("%.4f", w.Lon)
	meteoblueURL += "&lat=" + strLat + "&lon=" + strLon
	meteoblueURL += "&tz=" + configURL.TZ
	meteoblueURL += "&format=" + configURL.Format

	w.CurrentURL = meteoblueURL

	return meteoblueURL, nil
}
