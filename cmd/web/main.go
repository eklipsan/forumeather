package main

import (
	"net/http"
	// "github.com/eklipsan/forumeather/pkg/meteoblue"
)


func main() {
	app := NewApplication()
	// makeevka := meteoblue.NewLocation(48.0478,37.9258, "Makeevka")

	// result, err := makeevka.GetCurrentForecast()
	// if err != nil {
	// 	app.ErrorLog.Println(err)
	// 	return
	// }

	// app.InfoLog.Println(result)


	srv := http.Server{
		Addr: ":8080",
		Handler: app.routes(),
	}


	err := srv.ListenAndServe()
	if err != nil {
		app.ErrorLog.Println(err)
	}
}