package basic

import (
	"fmt"
	"regexp"
)

var ImageMatcher *regexp.Regexp
var packageVariable int = -54    // only visible in the basic package
var GlobalVariable int = 3451529 // globally visible

func init() {
	ImageMatcher = regexp.MustCompile(`[a-h]+[0-9]+\.jpg`)
	fmt.Printf("[init] ImageMatcher initialized to %s\n", ImageMatcher)
}

func Vars() {
	var count uint8
	var length uint16 = 10045
	name := "test"

	fmt.Printf("count = %d\nlength = %d\nname = %s\n", count, length, name)

	fmt.Printf("packageVariable = %d\n", packageVariable)
	fmt.Printf("GlobalVariable = %d\n", GlobalVariable)
}
