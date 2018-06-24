package views

import (
	"html/template"
	"log"
	"net/http"
)

// Render renders the template is the homepage
func Render(w http.ResponseWriter, viewPath string, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, err := template.ParseFiles("views/layouts/main.html", viewPath)
	if err != nil {
		log.Fatal(err)
		return
	}
	err2 := t.ExecuteTemplate(w, "layout", data)
	if err2 != nil {
		log.Fatal(err2)
		return
	}
}
