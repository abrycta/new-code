package main

import (
	"fmt"
	"os"
)

// usage:
// go run .04for2.go userinput
func main() {
	fmt.Println(os.Args) // represented as a slice of strings
	// first item in the slice is the executable
	// followed by the executable's arguments
	s, sep := "", ""                  // s is the resultant string, se
	for _, arg := range os.Args[1:] { // skip the executable / go file
		s += sep + arg // sep is the new word
		// append the empty string sep to the arg variable
		sep = "" // empty the string
	}
	fmt.Println(s)
}
