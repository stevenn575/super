package models

import (
	"testing"
)

func TestEmailValid(t *testing.T) {
	// u =: models.User{ "steve"}
	u := User{Email: "steve"}
	res := u.emailValid()
	if res {
		t.Error("Expected false, got ", res)
	}
	u.Email = "Steve@med.com"
	res2 := u.emailValid()
	if !res2 {
		t.Error("Expected true, got ", res2)
	}

}

func TestIsValid(t *testing.T) {
	var u User
	if u.isValid() {
		t.Error("Expected false if name and email not set, got", u.isValid())
	}
	u.Email = "Steve@Med.com"
	u.Name = "Steev"
	if !u.isValid() {
		t.Error("Expected true, got ", u.isValid())
	}
}
