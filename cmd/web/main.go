package main

import (
	"log"

	"html/template"
	"net/http"

	"github.com/go-martini/martini"
)

func main() {
	m := martini.Classic()
	m.Get("/", func(res http.ResponseWriter, req *http.Request) {
		t, err := template.ParseFiles("cmd/web/index.gtpl")
		if err != nil {
			log.Fatal(err)
		}
		t.Execute(res, nil)
	})

	m.Post("/results", func(r *http.Request) string {
		url := r.FormValue("url")
		return showPr(url)
	})
	m.Run()
}
