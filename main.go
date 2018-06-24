package main

import (
	"log"
	"net/http"
	"super/controllers"
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
	http.HandleFunc("/", controllers.HomeIndex)
	http.HandleFunc("/users", controllers.UsersIndex)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
