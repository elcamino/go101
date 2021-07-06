package basic

import (
	"fmt"
	"reflect"
)

func Types() {
	// Integers
	//
	width, length := 154, 3 // int
	fmt.Printf("width is a %s\n", reflect.TypeOf(width))

	var height int = 45          // int
	var items int8 = 125         // int8, int16, int32, int64
	var population uint = 476543 // uint, uint8, uint16, uint32, uint64

	// Floats
	//
	radius := 4.32 // float64
	fmt.Printf("radius is a %s\n", reflect.TypeOf(radius))
	var temperature float32 = 18.7 // float32, float64

	// Strings
	//
	city, state := "Nürnberg", "Bayern" // string

	// Boolean
	//
	inBayern := state == "Bayern"
	var letPass bool = false

	fmt.Println(width * length * height * int(items) * int(population))
	fmt.Println(radius * float64(temperature))
	fmt.Printf("%s, %s - pop. %d, in Bayern: %v\n", city, state, population, inBayern)
	fmt.Printf("pass: %v\n", letPass)

	// Maps
	//
	products := map[string]float32{
		"Coca Cola 1l":  1.9,
		"Pepsi Cola 1l": 1.87,
		"Dr. Pepper 1l": 1.87,
		"Fanta 1l":      0.99,
	}

	for k, v := range products {
		fmt.Printf("%s = €%.2f\n", k, v)
	}

	// Arrays
	//
	productCounts := []int32{
		2, 67, 234, 963323, 1253,
	}

	for i, v := range productCounts {
		fmt.Printf("%d: %d\n", i, v)
	}
}
