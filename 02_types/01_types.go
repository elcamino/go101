package types

import (
	"fmt"
	"time"

	"github.com/kr/pretty"
)

// new type "ExchangeRate"
// underlying type is map[string]float64
// map[string]float64 is a type litteral
type ExchangeRate map[string]float64

// new type "Birthdate"
// underlying type : time.Time (type Time from the time package)
type Birthdate time.Time

type Country struct {
	Name    string
	ISOCode string
}

// a new type Firstname
// underlying type : string
type Firstname string
type Lastname string

// a new type VATRate
// underlying type is float64
type VATRate float64

type Person struct {
	GivenName Firstname
	Surname   Lastname
	Born      Birthdate
	Country   *Country
	Website   string
}

func basic() {
	fmt.Println("\nBasic\n---------------------------------")
	p := Person{
		GivenName: "Jorge",
		Surname:   Lastname("Muller"),
		Born:      Birthdate(time.Now().Add(-1 * 6000 * 24 * time.Hour)),
		Country: &Country{
			Name:    "Bolivia",
			ISOCode: "BO",
		},
		Website: "http://www.dynamicback-end.org/innovate",
	}
	pretty.Println(p)

	bolivia := Country{
		"Bolivia",
		"BO",
	}
	pretty.Println(bolivia)
}

func embedded() {
	fmt.Println("\nEmbedded\n---------------------------------")
	type Book struct {
		Title string
		Year  uint
	}
	type Author struct {
		Person
		Books []Book
	}

	a := Author{
		Person: Person{
			GivenName: "Jorge",
			Surname:   Lastname("Muller"),
			Born:      Birthdate(time.Now().Add(-1 * 6000 * 24 * time.Hour)),
			Country: &Country{
				Name:    "Bolivia",
				ISOCode: "BO",
			},
			Website: "http://www.dynamicback-end.org/innovate",
		},
		Books: []Book{
			{Title: "Title A", Year: 1998},
			{Title: "Title B", Year: 2000},
		},
	}

	fmt.Printf("name: %s %s\n", a.GivenName, a.Person.Surname)
}

func Run() {
	basic()
	embedded()
	methods()
}
