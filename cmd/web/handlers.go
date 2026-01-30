package main

import (
	"html/template"
	"net/http"
	"path/filepath"
)


func (a Application) Home(w http.ResponseWriter, r *http.Request) {
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

	if err != nil {
		a.serverError(w, err)
	}
	err = ts.Execute(w, nil)
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