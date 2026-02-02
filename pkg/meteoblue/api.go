package meteoblue

import (
	"fmt"
	"io"
	"net/http"

	"github.com/eklipsan/forumeather/pkg/config"
)

type MeteobluePackage string

const (
	currentPackage MeteobluePackage = "current"
	basic15MinutePackage MeteobluePackage = "basic-15min"
	basic1HourPackage MeteobluePackage = "basic-1h"
	basic3HourPackage MeteobluePackage = "basic=3h"
	basicDayPackage MeteobluePackage = "basic-day"
)

var WeatherPictoCode = map[int64]string{
        1:  "Ясно и безоблачно",
        2:  "Ясно, небольшая облачность",
        3:  "Переменная облачность",
        4:  "Пасмурно",
        5:  "Туман",
        6:  "Пасмурно, дождь",
        7:  "Переменная облачность, временами легкий дождь",
        8:  "Легкий дождь, вероятны грозы",
        9:  "Пасмурно, снег",
        10: "Переменная облачность, временами снег",
        11: "Преимущественно облачно, дождь и снег",
        12: "Пасмурно, местами дождь",
        13: "Пасмурно, местами снег",
        14: "Преимущественно облачно, дождь",
        15: "Преимущественно облачно, снег",
        16: "Преимущественно облачно, местами дождь",
        17: "Преимущественно облачно, местами снег",
        18: "Не используется",
        19: "Не используется",
        20: "Преимущественно облачно",
        21: "Преимущественно ясно, возможны местные грозы",
        22: "Переменная облачность, возможны местные грозы",
        23: "Переменная облачность, возможны местные грозы и ливни",
        24: "Облачно с грозами и ливнями",
        25: "Преимущественно облачно, грозы и ливни",
    }


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
	currentUrl, err := w.createURL(currentPackage)
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
