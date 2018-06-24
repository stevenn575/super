package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// User is a user of Super
type User struct {
	Name    string
	Email   string
	TeamIDs []uint `json:"team_ids"`
	Teams   []Team
}

// GetUsers returns all records
func GetUsers() []User {
	raw, err := ioutil.ReadFile("./data/users.json")
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	var users []User
	json.Unmarshal(raw, &users)
	return users
}

func (u User) emailValid() bool {
	return strings.Contains(u.Email, "@") && strings.Contains(u.Email, ".")
}

func (u User) isValid() bool {
	if u.Email == "" {
		return false
	}
	if u.Name == "" {
		return false
	}
	return true
}

// GetTeams finds teams
func (u User) GetTeams() []Team {
	// Create a slice of size of the team_ids
	var teams []Team

	// Find each team
	for _, teamId := range u.TeamIDs {
		// Add it to slice
		log.Print(teamId)
		team, err := GetTeam(teamId)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
		teams = append(teams, team)
	}

	return teams
}
