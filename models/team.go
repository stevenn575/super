package models

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
)

// Team is a sports team
type Team struct {
	ID   uint
	Name string
	// League League
	Rank uint
}

// GetTeams returns all records
func GetTeams() []Team {
	var teams []Team
	var raw []byte
	var err error
	var filepath string

	if flag.Lookup("test.v") == nil {
		filepath = "./data/teams.json"
	} else {
		filepath = "../data/teams.json"
	}

	raw, err = ioutil.ReadFile(filepath)

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	json.Unmarshal(raw, &teams)
	return teams
}

// GetTeam finds a Team by ID
func GetTeam(id uint) (team Team, err error) {
	teams := GetTeams()
	for _, team := range teams {
		if team.ID == id {
			return team, nil
		}
	}
	return team, fmt.Errorf("Could not find team")
}
