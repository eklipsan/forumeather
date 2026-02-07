package main

import (
	"net/http"
)


func main() {
	app := NewApplication()

	srv := http.Server{
		Addr: app.Config.NetworkConfig.Port,
		Handler: app.routes(),
	}
	app.InfoLog.Printf("Создание сервера на порту %s", app.Config.NetworkConfig.Port)


	err := srv.ListenAndServe()
	if err != nil {
		app.ErrorLog.Println(err)
	}
	app.InfoLog.Println("Запуск сервера")
}