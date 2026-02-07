package main

import (
	"net/http"
	// "github.com/eklipsan/forumeather/pkg/meteoblue"
)


func main() {
	app := NewApplication()

	srv := http.Server{
		Addr: app.Config.NetworkConfig.Port,
		Handler: app.routes(),
	}


	err := srv.ListenAndServe()
	if err != nil {
		app.ErrorLog.Println(err)
	}
}