package main

import "net/http"

func (app *Application) routes() *http.ServeMux {
    mux := http.NewServeMux()
    mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/forum", app.ForumPage)

	fileServer := http.FileServer(NeuteredFileSystem{http.Dir("./ui/static/")})
	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

    return mux
}