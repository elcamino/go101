package basic

import "fmt"

func Names() {
	boostFactor := 3.2 // :-)

	boost_factor := 3.2             // :-(
	bf := 3.2                       // :-( ???
	BoostFactor := 3.2              // :-(
	boostFactorF := 3.2             // :-(
	boostReasonString := "because!" // :-(

	fmt.Printf("boost factor: %.2f, %s\n",
		(boostFactor+boost_factor+bf+BoostFactor+boostFactorF)/5.0,
		boostReasonString,
	)
}
