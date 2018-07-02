package models

import (
	"fmt"
	"log"
	"net/url"
	"strings"
	"super/system"
)

// User is a user of Super
type User struct {
	ID      uint
	Name    string
	Email   string
	TeamIDs []uint `json:"team_ids"`
	Teams   []Team
}

// GetUsers returns all records
func GetUsers() []User {
	var users []User

	db := system.GetDB()
	rows, err := db.Query("Select * from users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.ID, &user.Email, &user.Name)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
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

// CreateUser will create a new user
func CreateUser(data url.Values) (User, error) {
	var user User

	db := system.GetDB()

	statement := fmt.Sprintf("Insert into users(name, email) values('%v', '%v')", data.Get("user[name]"), "stevenn57@gmail.com")
	res, err := db.Exec(statement)
	if err != nil {
		return user, err
	}
	newID, _ := res.LastInsertId()
	log.Printf("Result of SQL execute is %v", newID)

	return user, err
}
