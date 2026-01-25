package meteoblue

import (
	"github.com/eklipsan/forumeather/pkg/config"
)



func CreateURL() (string, error) {
	var meteoblueURL string
	configURL, err := config.LoadMeteoblueConfig()

	if err != nil {
		return "", err
	}
	meteoblueURL = configURL.URL
	meteoblueURL += "?apikey=" + configURL.APIKey
	meteoblueURL += "&tz=" + configURL.TZ
	meteoblueURL += "&format=" + configURL.Format

	return meteoblueURL, nil
}