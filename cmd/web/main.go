package main

import (
	"net/http"
	// "github.com/eklipsan/forumeather/pkg/meteoblue"
)


func main() {
	app := NewApplication()

	srv := http.Server{
		Addr: ":8080",
		Handler: app.routes(),
	}


	err := srv.ListenAndServe()
	if err != nil {
		app.ErrorLog.Println(err)
	}
}