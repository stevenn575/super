package controllers

import (
	"html/template"
	"log"
	"net/http"
)

// HomeIndex is the homepage
func HomeIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, err := template.ParseFiles("views/layouts/main.html", "views/home/index.html")
	if err != nil {
		log.Fatal(err)
		return
	}
	err2 := t.ExecuteTemplate(w, "layout", nil)
	if err2 != nil {
		log.Fatal(err2)
		return
	}
}
