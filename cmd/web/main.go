package main

import (
	"fmt"
	"net/http"

	"github.com/eklipsan/forumeather/pkg/meteoblue"
)


func main() {
	makeevka := meteoblue.NewLocation(48.0478,37.9258, "Makeevka")

	result, err := makeevka.GetCurrentForecast()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)

	mux := http.NewServeMux()

	mux.HandleFunc("/", Home)

	fileServer := http.FileServer(NeuteredFileSystem{http.Dir("./ui/static/")})
	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))


	http.ListenAndServe(":8080", mux)
}
