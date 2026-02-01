package main

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/eklipsan/forumeather/pkg/meteoblue"
)



type ForumsCurrentWeather struct {
	ForumsInfo []ForumInfo
}

type ForumInfo struct {
	Name     string
	Place string
	Topics string
	Temperature float64
	Height int64
	Windspeed float64
	WeatherFromPictocode string
	Latitude float64
	Longitude float64
}


func (a Application) Home(w http.ResponseWriter, r *http.Request) {

	var forums ForumsCurrentWeather
	if r.URL.Path != "/" {
		a.notFound(w)
		return
	}
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		a.serverError(w, err)
	}

	dbForums, err := a.Database.GetAllForums()
	if err != nil {
		a.serverError(w, err)
	}

	for _, row := range dbForums {

		forumLocation := meteoblue.NewLocation(row.Latitude, row.Longitude, row.Name)
		forumCurrentForecast, err := forumLocation.GetCurrentForecast()
		if err != nil {
			a.serverError(w, err)
		}
		forumInfo := ForumInfo{
			Name: row.Name,
			Place: row.Place,
			Topics: row.Topics,
			Temperature: forumCurrentForecast.DataCurrent.Temperature,
			Height: forumCurrentForecast.Metadata.Height,
			Windspeed: forumCurrentForecast.DataCurrent.Windspeed,
			WeatherFromPictocode: meteoblue.WeatherPictoCode[forumCurrentForecast.DataCurrent.PictocodeDetailed],
			Latitude: row.Latitude,
			Longitude: row.Longitude,
		}
		forums.ForumsInfo = append(forums.ForumsInfo, forumInfo)
	}

	err = ts.Execute(w, forums)
	if err != nil {
		a.serverError(w, err)
	}

}

type NeuteredFileSystem struct {
    fs http.FileSystem
}

func (nfs NeuteredFileSystem) Open(path string) (http.File, error) {
	file, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}
	f, err := file.Stat()
	if err != nil {
		return nil, err
	}

	if f.IsDir() {
		index := filepath.Join(path, "index.html")
		dir, err := nfs.fs.Open(index)
		if err != nil {
			return nil, err
		} else {
			closeErr := dir.Close()
			return nil, closeErr
		}

	}
	return file, nil
}