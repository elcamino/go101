package logging

import (
	"fmt"
	"log"
	"os"
)

func Logging() {
	// the output can be set to any writer. os.Stderr is the default
	log.SetOutput(os.Stdout)

	// prints a line of text to the log
	log.Println("hello, this is a friendly log message!")

	// prints a formatted line of text to the log
	log.Printf("failed to load file after %d tries\n", 5)

	// add the file and line number to be logged with each message
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// prints a line of text to the log
	log.Println("hello, this is a friendly log message!")

	// prints a formatted line of text to the log
	log.Printf("failed to load file after %d tries\n", 5)

	// call a function that will panic and recover from it
	panicLog()
	log.Println("successfully recovered from panic")

	// log an error and exit the program
	log.Fatal("log.Fatal() always exits")
}

func panicLog() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("recovered from panic in panicLog()")
		}
	}()

	log.Panic("panicking now!")

	fmt.Println("after the panic")
}
