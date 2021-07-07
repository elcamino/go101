package arrays

import (
	"fmt"
	"math/rand"

	"github.com/kr/pretty"
)

func arrays() {
	fmt.Println("\nArrays\n---------------------------------")
	var arr [3]int
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}

	pretty.Println(arr)

	ids := make([]int, 5)
	for i := 0; i < len(ids); i++ {
		ids[i] = rand.Intn(100)
	}

	pretty.Println(ids)

	for i := 0; i < 5; i++ {
		ids = append(ids, rand.Intn(100))
	}
	pretty.Println(ids)
}

func twice(val *int) {
	*val *= 2
}

func pointers() {
	fmt.Println("\nPointers\n---------------------------------")
	var arr = make([]int, 5)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}
	pretty.Println(arr)
	for i := 0; i < len(arr); i++ {
		twice(&arr[i])
	}
	pretty.Println(arr)
}

func square(val int) int {
	return val * val
}

func slices() {
	fmt.Println("\nSlices\n---------------------------------")
	var arr [10]int // an array has a fixed length
	fmt.Printf("array len: %d, cap: %d\n", len(arr), cap(arr))

	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Printf("array len after setting values: %d, cap: %d\n", len(arr), cap(arr))

	pretty.Println(arr)

	// create a new slice with elems 2, 3 and 4 of arr
	//
	slice := arr[2:5] // a slice is a pointer to an array
	pretty.Println(slice)
	fmt.Printf("slice (2,5) len: %d, slice (2,5) cap: %d\n", len(slice), cap(slice))

	for i := 0; i < len(slice); i++ {
		slice[i] = square(slice[i])
	}
	pretty.Println(slice)
	pretty.Println(arr)

	intlist := []int{2, 3, 5, 7, 11, 13} // slice
	pretty.Println(intlist)

	madelist := make([]int, 10) // slice
	pretty.Println(madelist)

	// Length
	fmt.Printf("made slice len: %d, made slice cap: %d\n", len(madelist), cap(madelist))

	// append to a slice
	ids := make([]string, 0)
	for i := 0; i < 10; i++ {
		ids = append(ids, fmt.Sprintf("%d", rand.Uint32()))
	}
	pretty.Println(ids)
}

func Run() {
	arrays()
	slices()
	pointers()
}
