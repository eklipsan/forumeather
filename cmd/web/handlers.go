package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)


func Home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ts.Execute(w, nil)

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