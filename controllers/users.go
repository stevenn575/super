package controllers

import (
	"html/template"
	"log"
	"net/http"
	"super/models"
)

// UsersIndex lists all users
func UsersIndex(w http.ResponseWriter, r *http.Request) {
	u := models.GetUsers()

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, err := template.ParseFiles("views/layouts/main.html", "views/users/index.html")
	if err != nil {
		log.Fatal(err)
		return
	}
	err2 := t.ExecuteTemplate(w, "layout", u)
	if err2 != nil {
		log.Fatal(err2)
		return
	}
}
