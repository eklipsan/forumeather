package main

import (
	"net/http"

	"github.com/eklipsan/forumeather/pkg/meteoblue"
)


func main() {
	app := NewApplication()
	makeevka := meteoblue.NewLocation(48.0478,37.9258, "Makeevka")

	result, err := makeevka.GetCurrentForecast()
	if err != nil {
		app.ErrorLog.Println(err)
		return
	}

	app.InfoLog.Println(result)

	mux := http.NewServeMux()

	mux.HandleFunc("/", app.Home)

	fileServer := http.FileServer(NeuteredFileSystem{http.Dir("./ui/static/")})
	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))


	http.ListenAndServe(":8080", mux)
}
