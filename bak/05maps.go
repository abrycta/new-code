package main

import "fmt"

func main() {
	//ages := make(map[string]int) // mapping from strings to ints
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	/*
		This is equivalent to
		ages := make(map[string]int)
		ages["alice"] = 31
		ages["charlie"] = 34
	*/

	//Map elements are accessed through the usual subscript notation:
	ages["alice"] = 32
	fmt.Println("alice = ", ages["alice"]) // "32"

	//removed with the built-in function delete:
	delete(ages, "alice") // remove element ages["alice"]

	/*
		All of these operations are safe even if the element isn’t in the map
	*/
	ages["bob"] = ages["bob"] + 1
	/*
		or
		ages["bob"] += 1
		or
		ages["bob"]++
	*/
	fmt.Println("bob = ", ages["bob"])
	/* a map element is not a variable, and we cannot take its address:

	_=&ages["bob"] // compile error: cannot take address of map element

	*/

}
