package basic

import "fmt"

func xyz() string { return "xyz " }

func Consts() {
	// default type is bool
	const isOpen = true
	// default type is rune (alias for int32, exported)
	const MyRune = 'r'
	// default type is int
	const occupancyLimit = 12
	// default type is float64 (exported)
	const VATRate = 29.87
	// default type is complex128
	const complexNumber = 1 + 2i
	// default type is string
	const hotelName = "Gopher Hotel"

	// const can only be a primitive type
	//const doesntWork = xyz()
	//const mapDoesntWork = map[string]string{}

	// consts are unlimited
	const profit = 9223372036854775808
	// var profit2 int64 = profit

	const (
		HOTEL = iota
		BNB
		ROOM
		HOUSE
		CHALET
		VILLA
		BOAT
		TREEHOUSE
	)
	const (
		URBAN = iota
		RURAL
	)

	fmt.Printf("HOTEL = %d\n", HOTEL)
	fmt.Printf("BNB = %d\n", BNB)
	fmt.Printf("ROOM = %d\n", ROOM)
	fmt.Printf("CHALET = %d\n", CHALET)
	fmt.Printf("VILLA = %d\n", VILLA)
	fmt.Printf("BOAT = %d\n", BOAT)
	fmt.Printf("TREEHOUSE = %d\n", TREEHOUSE)
	fmt.Printf("URBAN = %d\n", URBAN)
	fmt.Printf("RURAL = %d\n", RURAL)

}
