package funcs

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var pricePerNight = 100.0

func computePrice(rate float32, nights int) float32 {
	return rate * float32(nights)
}

func defaultPricePerNight() float64 {
	return pricePerNight
}

func setDefaultPricePerNight(price float64) float64 {
	pricePerNight = price
	return pricePerNight
}

func scope() {
	fmt.Printf("default price per night: $%.2f\n", defaultPricePerNight())

	pricePerNight := 251.34
	nights := 5

	fmt.Printf("total price for %d nights at $%.2f: $%.2f\n", nights, pricePerNight, computePrice(float32(pricePerNight), nights))
	fmt.Printf("default price per night: $%.2f\n", defaultPricePerNight())
	setDefaultPricePerNight(200.0)
	fmt.Printf("default price per night: $%.2f\n", defaultPricePerNight())
}

func computePriceNamedReturnValue(rate float32, nights int) (price float32) {
	price = rate * float32(nights) * 2
	return
}

func namedReturn() {
	pricePerNight := 251.34
	nights := 5

	fmt.Printf("total price for %d nights at $%.2f: $%.2f\n", nights, pricePerNight, computePriceNamedReturnValue(float32(pricePerNight), nights))
}

// BookRoom books a number of persons into a room (exported)
func BookRoom(roomNumber int, persons ...string) {
	for _, p := range persons {
		fmt.Printf("booking room #%d for %s\n", roomNumber, p)
	}
}

func variadic() {
	BookRoom(14, "Homer", "Marge", "Bart", "Lisa")
}

type list []int

func (l list) Each(callback func(val int)) {
	for _, v := range l {
		callback(v)
	}
}

func anonymous() {
	l := list{
		1, 2, 3, 7, 8, 9,
	}

	l.Each(func(v int) {
		fmt.Printf("%d * %d = %d\n", v, v, v*v)
	})

	callback := func(v int) {
		fmt.Printf("[callback] %d * %d = %d\n", v, v, v*v)
	}

	l.Each(callback)
}

func closures() {
	name := "Bart"
	now := time.Now()

	// a closure has access to all variables in the context from which it is called
	func() {
		fmt.Printf("Hey %s! It's %s...time to get up!\n", name, now.Format("15:04"))
	}()
}

func defers() {
	fmt.Println("doing something func...")

	defer func() { fmt.Println("defer func 1") }()
	defer func() { fmt.Println("defer func 2") }()

	fh, err := os.Open("data/rand.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fh.Close() // close the filehandle before leaving the function

	data, err := ioutil.ReadAll(fh)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

}

func Run() {
	scope()
	namedReturn()
	variadic()
	anonymous()
	closures()
	defers()
}
