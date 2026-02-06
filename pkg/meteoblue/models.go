package meteoblue

// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    root, err := UnmarshalRoot(bytes)
//    bytes, err = root.Marshal()

import "encoding/json"

func UnmarshalCurrentForecast(data []byte) (CurrentForecast, error) {
	var r CurrentForecast
	err := json.Unmarshal(data, &r)
	return r, err
}

// структура для текущей погоды
type CurrentForecast struct {
	Metadata    metadataCurrent    `json:"metadata"`
	Units       unitsCurrent       `json:"units"`
	DataCurrent dataCurrent `json:"data_current"`
}

type dataCurrent struct {
	Time              string      `json:"time"`
	Isobserveddata    int64       `json:"isobserveddata"`
	Metarid           any 		  `json:"metarid"`
	Isdaylight        int64       `json:"isdaylight"`
	Windspeed         float64     `json:"windspeed"`
	Zenithangle       float64     `json:"zenithangle"`
	PictocodeDetailed int64       `json:"pictocode_detailed"`
	Pictocode         int64       `json:"pictocode"`
	Temperature       float64     `json:"temperature"`
}

type metadataCurrent struct {
	ModelrunUpdatetimeUTC string  `json:"modelrun_updatetime_utc"`
	Name                  string  `json:"name"`
	Height                int64   `json:"height"`
	TimezoneAbbrevation   string  `json:"timezone_abbrevation"`
	Latitude              float64 `json:"latitude"`
	ModelrunUTC           string  `json:"modelrun_utc"`
	Longitude             float64 `json:"longitude"`
	UTCTimeoffset         float64   `json:"utc_timeoffset"`
	GenerationTimeMS      float64 `json:"generation_time_ms"`
}

type unitsCurrent struct {
	Temperature string `json:"temperature"`
	Time        string `json:"time"`
	Windspeed   string `json:"windspeed"`
}




func UnmarshalTrendDayForecast(data []byte) (TrendDayForecast, error) {
	var r TrendDayForecast
	err := json.Unmarshal(data, &r)
	return r, err
}

// структура для 14-дневного прогноза
type TrendDayForecast struct {
	Metadata metadataTrendDay `json:"metadata"`
	Units    unitsTrendDay    `json:"units"`
	TrendDay trendDay `json:"trend_day"`
}

type metadataTrendDay struct {
	ModelrunUpdatetimeUTC string  `json:"modelrun_updatetime_utc"`
	Name                  string  `json:"name"`
	Height                int64   `json:"height"`
	TimezoneAbbrevation   string  `json:"timezone_abbrevation"`
	Latitude              float64 `json:"latitude"`
	ModelrunUTC           string  `json:"modelrun_utc"`
	Longitude             float64 `json:"longitude"`
	UTCTimeoffset         float64   `json:"utc_timeoffset"`
	GenerationTimeMS      float64 `json:"generation_time_ms"`
}

type trendDay struct {
	Time                           []string  `json:"time"`
	TotalcloudcoverMax             []int64   `json:"totalcloudcover_max"`
	ExtraterrestrialradiationTotal []int64   `json:"extraterrestrialradiation_total"`
	TotalcloudcoverMin             []int64   `json:"totalcloudcover_min"`
	Predictability                 []int64   `json:"predictability"`
	Precipitation                  []float64 `json:"precipitation"`
	TemperatureMax                 []float64 `json:"temperature_max"`
	SealevelpressureMean           []int64   `json:"sealevelpressure_mean"`
	SealevelpressureMin            []int64   `json:"sealevelpressure_min"`
	WindspeedMean                  []float64 `json:"windspeed_mean"`
	Pictocode                      []int64   `json:"pictocode"`
	Snowfraction                   []float64   `json:"snowfraction"`
	RelativehumidityMax            []int64   `json:"relativehumidity_max"`
	TotalcloudcoverMean            []int64   `json:"totalcloudcover_mean"`
	TemperatureMin                 []float64 `json:"temperature_min"`
	Winddirection                  []int64   `json:"winddirection"`
	PrecipitationSpread            []float64 `json:"precipitation_spread"`
	RelativehumidityMin            []int64   `json:"relativehumidity_min"`
	WindspeedMin                   []float64 `json:"windspeed_min"`
	WindspeedSpread                []float64 `json:"windspeed_spread"`
	PrecipitationProbability       []int64   `json:"precipitation_probability"`
	TotalcloudcoverSpread          []int64   `json:"totalcloudcover_spread"`
	RelativehumidityMean           []int64   `json:"relativehumidity_mean"`
	TemperatureMean                []float64 `json:"temperature_mean"`
	SealevelpressureMax            []int64   `json:"sealevelpressure_max"`
	GhiTotal                       []int64   `json:"ghi_total"`
	TemperatureSpread              []float64 `json:"temperature_spread"`
	PredictabilityClass            []int64   `json:"predictability_class"`
	WindspeedMax                   []float64 `json:"windspeed_max"`
}

type unitsTrendDay struct {
	Predictability           string `json:"predictability"`
	Precipitation            string `json:"precipitation"`
	Windspeed                string `json:"windspeed"`
	Cloudcover               string `json:"cloudcover"`
	PrecipitationProbability string `json:"precipitation_probability"`
	Relativehumidity         string `json:"relativehumidity"`
	Radiationtotal           string `json:"radiationtotal"`
	Time                     string `json:"time"`
	Temperature              string `json:"temperature"`
	Pressure                 string `json:"pressure"`
	Winddirection            string `json:"winddirection"`
}