package maps

import "fmt"

func Maps() {
	requests := map[string]int{
		"Chrome":  1000,
		"Firefox": 400,
		"Opera":   30,
	}

	requests["Lynx"] = 5

	fmt.Println("\nrequests\n---------------------------------------------------")
	for k, v := range requests {
		fmt.Printf("%s = %d\n", k, v)
	}

	// merge data into the requests map from another requests map
	requests2 := make(map[string]int)
	requests2["Chrome"] = 451
	requests2["Firefox"] = 247
	requests2["Opera"] = 19
	requests2["SeaMonkey"] = 10

	for k, v := range requests2 {
		if v2, exists := requests[k]; exists {
			requests[k] = v2 + v
		} else {
			requests[k] = v
		}
	}

	fmt.Println("\nrequests after merge\n---------------------------------------------------")
	for k, v := range requests {
		fmt.Printf("%s = %d\n", k, v)
	}
}
