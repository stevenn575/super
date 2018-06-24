package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

// User is a user of Super
type User struct {
	Name  string
	Email string
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
