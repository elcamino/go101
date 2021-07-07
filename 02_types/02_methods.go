package types

import (
	"fmt"

	"github.com/kr/pretty"
)

type User struct {
	FirstName string
	LastName  string
	Email     string
}

// value receiver: everything in u gets copied before the method is called.
// you can't modify any property of the original receiver
func (u User) SetFirstNameNotWorking(name string) {
	u.FirstName = name
}

// pointer receiver: can modify the underlying data
func (u *User) SetFirstName(name string) {
	u.FirstName = name
}

func (u *User) Name() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

func (u *User) NameAndEmail() (string, string) {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName), u.Email
}

// private method: only visible within User
func (u *User) buildFullEmailAddress() string {
	return fmt.Sprintf("%s <%s>", u.Name(), u.Email)
}

func (u *User) FullEmail() string {
	return u.buildFullEmailAddress()
}

func methods() {
	fmt.Println("\nMethods\n---------------------------------")

	user := User{
		FirstName: "Hayley",
		LastName:  "Ferry",
		Email:     "savannahpagac@murray.com",
	}

	pretty.Println(user)
	user.SetFirstNameNotWorking("Gerard")
	pretty.Println(user)
	user.SetFirstName("Gerard")
	pretty.Println(user)

	fmt.Printf("name: %s\n", user.Name())

	uname, uemail := user.NameAndEmail()
	fmt.Printf("%s <%s>\n", uname, uemail)
	fmt.Println(user.FullEmail())
}
