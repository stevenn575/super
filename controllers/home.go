package controllers

import (
	"net/http"
	"super/views"
)

// HomeIndex is the homepage
func HomeIndex(w http.ResponseWriter, r *http.Request) {
	views.Render(w, "views/home/index.html", nil)
}
