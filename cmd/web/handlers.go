package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/eklipsan/forumeather/pkg/meteoblue"
	"github.com/eklipsan/forumeather/pkg/storage"
)



type ForumsCurrentWeather struct {
	ForumsInfo []ForumInfoCurrent
	SearchQuery string
	TotalPages []int
	CurrentPage int
}

type ForumInfoCurrent struct {
	ID int64
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
		forumInfo := ForumInfoCurrent{
			ID: row.ID,
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

type ForumTrendWeather struct {
	ForumInfo ForumInfoCurrent
	ForumTrend []ForumInfoTrend
}

type ForumInfoTrend struct {
	Time string
	WeatherFromPictocode string
	TemperatureMax float64
	TemperatureMin float64
	WindSpeedMean float64
	PrecipitationProbability int64
	Predictability int64
}



func (a Application) ForumPage(w http.ResponseWriter, r *http.Request) {
	var ForumDaysTrend []ForumInfoTrend
	files := []string{
		"./ui/html/forum.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		a.serverError(w, err)
	}

	forum_id_str := r.URL.Query().Get("id")
	forum_id_int, _ := strconv.Atoi(forum_id_str)
	a.InfoLog.Println("id:", forum_id_int)


	forumDB, err := a.Database.GetForum(forum_id_int)
	if err != nil {
		a.serverError(w, err)
	}
	a.InfoLog.Println(forumDB)


	forumLocation := meteoblue.NewLocation(forumDB.Latitude, forumDB.Longitude, forumDB.Name)
	trendDayForecast, err := forumLocation.GetTrendDayForecast()
	if err != nil {
		a.serverError(w, err)
	}

	ForumInfoCurrent := ForumInfoCurrent{
		Name: forumDB.Name,
		Place: forumDB.Place,
		Topics: forumDB.Topics,
		Temperature: 0.0,
		Height: 0,
	}

	for dayIndex := 0; dayIndex < len(trendDayForecast.TrendDay.Time) - 4; dayIndex ++ {
		var DayTrend ForumInfoTrend
		// "2006-01-02 15:04:05" "YYYY-MM-DD hh:mm",
		ParsedTime, err := time.Parse("2006-01-02 15:04", trendDayForecast.TrendDay.Time[dayIndex])
		if err != nil {
			a.ErrorLog.Println(err)
		}

		DayTrend.Time = GetTimeLabel(ParsedTime)
		DayTrend.WeatherFromPictocode = meteoblue.WeatherPictoCode[trendDayForecast.TrendDay.Pictocode[dayIndex]]
		DayTrend.TemperatureMax = trendDayForecast.TrendDay.TemperatureMax[dayIndex]
		DayTrend.TemperatureMin = trendDayForecast.TrendDay.TemperatureMin[dayIndex]
		DayTrend.WindSpeedMean = trendDayForecast.TrendDay.WindspeedMean[dayIndex]
		DayTrend.PrecipitationProbability = trendDayForecast.TrendDay.PrecipitationProbability[dayIndex]
		DayTrend.Predictability = trendDayForecast.TrendDay.Predictability[dayIndex]

		ForumDaysTrend = append(ForumDaysTrend, DayTrend)
	}

	ForumPage := ForumTrendWeather{
		ForumInfo: ForumInfoCurrent,
		ForumTrend: ForumDaysTrend,
	}

	err = ts.Execute(w, ForumPage)
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