package types

import (
	"fmt"
	"math"
	"reflect"
	"time"
)

// Defining the geometry interface
//
type geometry interface {
	area() float64
	perim() float64
}

// rect represents a rectangle and implements the geometry interface
//
type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// circle represents a, you guessed it: circle and also implements the geometry interface
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// printMeasurements prints the area and perimeter of a geometry interface
//
func printMeasurements(g geometry) {
	fmt.Printf("%s = %v - area: %f - perimeter: %f\n", reflect.TypeOf(g), g, g.area(), g.perim())
}

func emptyInterface() {
	// the interface{} slot can contain any type
	m := map[int]interface{}{
		0: 1024,
		1: "number one",
		2: time.Now(),
		3: func(name string) { fmt.Printf("Hey %s, you're fourth!\n", name) },
	}

	fmt.Printf("\nEmpty interfaces\n---------------------------------------------------------\n")
	for k, v := range m {
		fmt.Printf("%d = %s\n", k, reflect.TypeOf(v))
	}

	// cast interface{} to the actual type (time.Time)
	t := m[2].(time.Time)
	fmt.Println(t)

	// call the stored function
	m[3].(func(string))("Joe")

	// casting to the wrong type results in a panic
	i := m[1].(int)
	fmt.Println(i)
}

func interfaceDemo() {
	fmt.Printf("\nInterfaces\n---------------------------------------------------------\n")

	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	printMeasurements(r)
	printMeasurements(c)
}

func Interfaces() {
	interfaceDemo()
	emptyInterface()
}
