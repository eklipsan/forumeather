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

func (r *CurrentForecast) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CurrentForecast struct {
	Metadata    metadata    `json:"metadata"`
	Units       units       `json:"units"`
	DataCurrent dataCurrent `json:"data_current"`
}

type dataCurrent struct {
	Time              string      `json:"time"`
	Isobserveddata    int64       `json:"isobserveddata"`
	Metarid           interface{} `json:"metarid"`
	Isdaylight        int64       `json:"isdaylight"`
	Windspeed         float64     `json:"windspeed"`
	Zenithangle       float64     `json:"zenithangle"`
	PictocodeDetailed int64       `json:"pictocode_detailed"`
	Pictocode         int64       `json:"pictocode"`
	Temperature       float64     `json:"temperature"`
}

type metadata struct {
	ModelrunUpdatetimeUTC string  `json:"modelrun_updatetime_utc"`
	Name                  string  `json:"name"`
	Height                int64   `json:"height"`
	TimezoneAbbrevation   string  `json:"timezone_abbrevation"`
	Latitude              float64 `json:"latitude"`
	ModelrunUTC           string  `json:"modelrun_utc"`
	Longitude             float64 `json:"longitude"`
	UTCTimeoffset         int64   `json:"utc_timeoffset"`
	GenerationTimeMS      float64 `json:"generation_time_ms"`
}

type units struct {
	Temperature string `json:"temperature"`
	Time        string `json:"time"`
	Windspeed   string `json:"windspeed"`
}
