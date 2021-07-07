package timedemo

import (
	"fmt"
	"log"
	"time"
)

func Time() {
	fmt.Println("what time is it?\n-------------------------------------------------------------------------------")
	now := time.Now()

	// Add adds time.Duration to time.Time and returns time.Time
	oneHourAgo := now.Add(-1 * (time.Hour + 4*time.Minute + 31*time.Second + 150*time.Millisecond + 5*time.Microsecond + 126*time.Nanosecond))

	// Sub subtracts two times and returns time.Duration
	timeDiff := now.Sub(oneHourAgo)

	fmt.Printf("now:         %s\n", now)
	fmt.Printf("unix time:   %d\n", now.Unix())
	fmt.Printf("an hour ago: %s\n", oneHourAgo)
	fmt.Printf("difference:  %s\n", timeDiff)

	fmt.Println("\nparsing a time, claculating the difference and truncating\n-------------------------------------------------------------------------------")

	// formatting and parsing a time is a bit awkward in Go.
	// Please refer to the documentation at
	// https://golang.org/pkg/time/#pkg-constants
	//
	openingTime, err := time.Parse(time.RFC3339, fmt.Sprintf("%sT18:30:00+02:00", now.Format("2006-01-02")))
	if err != nil {
		log.Fatal(err)
	}

	timeUntilOpening := openingTime.Sub(now).Truncate(30 * time.Minute)

	fmt.Printf("%s until the opening time (%s)\n", timeUntilOpening, openingTime.Format("15:04"))

	fmt.Println("\ntickers\n-------------------------------------------------------------------------------")

	ticker := time.NewTicker(time.Second)

	// make sure to stop the ticker or else it will keep running and leak memory
	defer ticker.Stop()

	endTime := time.Now().Add(5 * time.Second)

	for range ticker.C {
		// stop after at most 5 seconds
		if time.Now().After(endTime) {
			fmt.Println("exiting ticker loop")
			break
		}

		fmt.Println(time.Now())
	}
}
