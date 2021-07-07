package types

import "testing"

func TestUserName(t *testing.T) {
	u := User{}
	carlsName := "Carl"
	ronaldsName := "Ronald"

	// verify SetFirstName is working
	u.SetFirstName(carlsName)

	if u.FirstName != carlsName {
		t.Errorf("SetFirstName didn't set FirstName to '%s' but to '%s'", carlsName, u.FirstName)
	}

	// verify SetFirstNameNotWorking doesn't modify FirstName
	// since the receiver is a copy of u
	u.SetFirstNameNotWorking(ronaldsName)

	if u.FirstName != carlsName {
		t.Errorf("SetFirstNameNotWorking set FirstName to '%s' but should have left it '%s'", u.FirstName, carlsName)
	}
}
