package main

import "fmt"

var l int = 24 // full variable declaration syntax at package level
const CONSTANTINT = 666

func main() {
	// alternative forms
	// var i = 0
	// var i int = 0
	var i int // declare a variable i of type int
	// i is zero, no such thing as an uninitialized variable
	// equal to i := 0

	j, k := 10, 11 // j initialized to 10
	// k initialized to 11
	// j := 20 is invalid, := notation is
	// only valid for declarations, not
	// method assignments

	// These are all declarations, not assignments:
	var l, m, n int                 // l, m, n = 0
	var b, f, s = true, 2.3, "four" // var is NEVER USED with the := notation
	// b = true, f = 2.3, s = "four"

	fmt.Printf("%v, %T\n", i, i)           //prints value of variable and variable type
	fmt.Println("j=", j, "k=", k, "l=", l) //variables must always be used in Go, otherwise you get a compile-time error.
	fmt.Println(CONSTANTINT)
	fmt.Println("m=", m, "n=", n)
	fmt.Println("b=", b, "f=", f, "s=", s)
}
