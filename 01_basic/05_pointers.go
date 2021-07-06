package basic

import "fmt"

func pointerBasics() {
	var ptr *int

	fmt.Printf("ptr is nil: %v\n", ptr == nil)

	var i int = 19

	// assign the address of i to ptr
	ptr = &i
	fmt.Printf("ptr is nil: %v, ptr (addr %d) is indeed %d\n", ptr == nil, ptr, *ptr)

	// dereference the pointer before working with it
	*ptr++
	fmt.Printf("ptr is nil: %v, ptr (addr %d) is indeed %d\n", ptr == nil, ptr, *ptr)

	*ptr = 22
	fmt.Printf("ptr is nil: %v, ptr (addr %d) is indeed %d\n", ptr == nil, ptr, *ptr)
}

func Pointers() {
	pointerBasics()
}
