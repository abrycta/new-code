package main

import (
	"fmt"
)

func main() {

	//Array initialization
	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May",
		6: "June", 7: "July", 8: "August", 9: "September",
		10: "October", 11: "November", 12: "December"}

	//Slice initialization
	/*
		The slice operator s[i:j], where 0 ≤ i ≤ j ≤ cap(s), creates a new slice that refers to elements
		i through j-1 of the sequence s, which may be an array variable, a pointer to an array, or
		another slice. The resulting slice has j-i elements.
	*/
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)     // ["April" "May" "June"]
	fmt.Println(summer) // ["June" "July" "August"]

	fmt.Println()

	// June is included in each and is the sole out put of this (inefficient) test for common elements:
	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s) // June appears in both
			}
		}
	}
	fmt.Println()

	/*
		Slicing beyond cap(s) causes a panic, but slicing beyond len(s) extends the slice, so the
		result may be longer than the original:
	*/
	//fmt.Println(summer[:20]) // panic: out of range
	endlessSummer := summer[:5] // extend a slice (within capacity)
	fmt.Println(endlessSummer)  // "[June July August September October]"
	fmt.Println()

	// Append function
	// The built-in append function appends items to slices
	var runes []rune
	for _, r := range "Hello, World" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes) // "['H' 'e' 'l' 'l' 'o' ',' ' ' 'W' 'o' 'r' 'l' 'd']"
}
