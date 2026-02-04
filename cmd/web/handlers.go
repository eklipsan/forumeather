package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/eklipsan/forumeather/pkg/meteoblue"
	"github.com/eklipsan/forumeather/pkg/storage"
)



type ForumsCurrentWeather struct {
	ForumsInfo []ForumInfo
	SearchQuery string
	TotalPages []int
	CurrentPage int
}

type ForumInfo struct {
	Name     string
	Place string
	Topics string
	Temperature float64
	Height int64
	Windspeed float64
	WeatherFromPictocode string
}

type SearchFilter struct {
	Query string
}


func (a Application) Home(w http.ResponseWriter, r *http.Request) {
	var (
		dbForums []storage.Forum
		forums ForumsCurrentWeather

		SliceTotalPage []int
		pageSize = 6
	)
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


	searchQuery := r.URL.Query().Get("search")
	currentPage := r.URL.Query().Get("page")
	if currentPage == "" { currentPage = "1"}

	currentPageInt, _ := strconv.Atoi(currentPage)

	if searchQuery == "" {
		dbForums, err = a.Database.GetAllForums()
		if err != nil {
			a.serverError(w, err)
		}
	} else {
		dbForums, err = a.Database.SearchForums(searchQuery)
		if err != nil {
			a.serverError(w, err)
		}

	}


	totalItems := len(dbForums)
    totalPages := (totalItems + pageSize - 1) / pageSize

	start := (currentPageInt - 1) * pageSize
    end := start + pageSize

	if end > totalItems { end = totalItems }

	if start < totalItems {
        dbForums =  dbForums[start:end]
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
		}
		forums.ForumsInfo = append(forums.ForumsInfo, forumInfo)
	}


	forums.SearchQuery = searchQuery

	for i := 1; i <= totalPages; i++ {
		SliceTotalPage = append(SliceTotalPage, i)
	}
	forums.TotalPages = SliceTotalPage
	forums.CurrentPage = currentPageInt

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