package controllers

import (
	"net/http"
	"super/models"
	"super/views"
)

// UsersIndex lists all users
func UsersIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		models.CreateUser(r.Form)
	}
	users := models.GetUsers()
	for i, user := range users {
		users[i].Teams = user.GetTeams()
	}
	views.Render(w, "views/users/index.html", users)
}
