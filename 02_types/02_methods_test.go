package types

import "testing"

func TestUserName(t *testing.T) {
	u := User{}

	// verify SetFirstName is working
	u.SetFirstName("Carl")

	if u.FirstName != "Carl" {
		t.Errorf("SetFirstName didn't set FirstName to 'Carl' but to '%s'", u.FirstName)
	}

	// verify SetFirstNameNotWorking doesn't modify FirstName
	u.SetFirstNameNotWorking("Ronald")

	if u.FirstName != "Carl" {
		t.Errorf("SetFirstNameNotWorking set FirstName to '%s'", u.FirstName)
	}
}
