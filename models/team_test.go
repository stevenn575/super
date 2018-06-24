package models

import (
	"testing"
)

func TestGetTeams(t *testing.T) {
	res := GetTeams()
	if res == nil {
		t.Error("Should return teams, got ", res)
	}
}
