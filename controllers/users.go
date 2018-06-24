package controllers

import (
	"net/http"
	"super/models"
	"super/views"
)

// UsersIndex lists all users
func UsersIndex(w http.ResponseWriter, r *http.Request) {
	u := models.GetUsers()
	views.Render(w, "views/users/index.html", u)

}
