package control

import (
	"fmt"
	"math/rand"
	"time"
)

func Run() {

	// create two random numbers with a being larger than b
	//
	rand.Seed(time.Now().UnixNano())
	rnd := rand.Intn(80)
	a := rnd + 20
	b := rand.Intn(rnd)

	// Comparisions
	//
	fmt.Printf("%d == %d: %v\n", a, b, a == b)
	fmt.Printf("%d == %d: %v\n", a, a, a == a)
	fmt.Printf("%d != %d: %v\n", a, b, a != b)
	fmt.Printf("%d > %d: %v\n", a, b, a > b)
	fmt.Printf("%d < %d: %v\n", a, b, a < b)
	fmt.Printf("%d <= %d: %v\n", a, b, a <= b)
	fmt.Printf("%d <= %d: %v\n", a, a, a <= a)
	fmt.Printf("%d >= %d: %v\n", a, a, a >= a)

	// if / then / else
	//
	if a > b {
		fmt.Printf("%d > %d\n", a, b)
	} else if a < b {
		fmt.Printf("%d < %d\n", a, b)
	} else {
		fmt.Printf("%d == %d\n", a, b)
	}

	// switch
	//
	switch {
	case a > 66:
		fmt.Println("a: *********")
	case a > 33:
		fmt.Println("a: ******")
	default:
		fmt.Println("a: ***")
	}

	switch a {
	case 0:
		fmt.Println("ZEROOOOO!")
	default:
		fmt.Println("not zero")
	}

	const emailsToSend = 3
	emailsSent := 0

	// for loop
	//
	for emailsSent < emailsToSend {
		emailsSent++
		fmt.Printf("sent email %d of %d\n", emailsSent, emailsToSend)
	}

	for i := 0; i < emailsToSend; i++ {
		fmt.Printf("sending email %d of %d\n", i+1, emailsToSend)
	}

	emailsSent = 0
	for {
		if emailsSent >= emailsToSend {
			break
		}

		fmt.Printf("sending email %d of %d\n", emailsSent+1, emailsToSend)
		emailsSent++
	}

	// breaking out of stacked for loops
OUTER:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i > 1 {
				break OUTER
			}
			fmt.Printf("i: %d, j: %d\n", i, j)
			if j > 2 {
				break
			}

		}
	}
}
