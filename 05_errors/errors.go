package errors

import (
	"fmt"
	"io/ioutil"
	"log"
)

func readFile(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	// an error has occurred: return no data and the error
	if err != nil {
		return nil, err
	}

	// no error has occurred: return the data read and no error
	return data, nil
}

func Errors() {
	fmt.Println("\nreading an existing file\n---------------------------------------------------")

	data, err := readFile("data/rand.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))

	fmt.Println("\nreading a file that doesn't exist\n---------------------------------------------------")

	data, err = readFile("data/rand2.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}
