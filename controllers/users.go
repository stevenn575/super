package controllers

import (
	"log"
	"net/http"
	"strconv"
	"super/models"
	"super/views"

	"github.com/gorilla/mux"
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

// UsersDelete deletes a user
func UsersDelete(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	models.DeleteUser(id)
	http.Redirect(w, r, "/users", 302)
}
