package main

import (
	"html/template"
	"log"
	"net/http"
	"super/models"
)

// UserLeague is a group of UserTeams
type UserLeague struct {
	UserTeams []UserTeam
}

// UserTeam is is the teams that a user picks
type UserTeam struct {
	Teams []models.Team
}

// League has many teams
type League struct {
	Name string
}

func main() {
	log.Print("Running")

	// Make a webpage
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	u := models.GetUsers()

	log.Print(u)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, err := template.ParseFiles("views/layout.html", "views/users.html")
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
