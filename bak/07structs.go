package main

import "fmt"

// Declare a Person object
type Person struct {
	name string
	age  int
}

// Function is declared as an initializer
// of a person struct
func newPerson(name string) *Person {
	// new(T) returns the pointer of a Type T
	// new uses the default values
	return new(Person)
}

// Go does not support method overloading
// uncommenting this method will throw a
// compile-time error: 'newPerson'
// redeclared in this package
/*
func newPerson(name string, age int) {

}

*/

func main() {
	// Declare a named Person instance
	bobPerson := newPerson("Bob")
	fmt.Println(bobPerson)

	// Print out the values of an anonymous Person instance
	// use the shorthand initializer:
	// Type{property: value}
	fmt.Println(Person{name: "Alice", age: 30})

	// the 'new' keyword returns a pointer to
	// newly created Person instance
	// use *variable notation to access the value
	// stored at the address of the pointer
	var fredPerson Person = *new(Person)
	fredPerson.name = "Fred"
	fredPerson.age = 51
	fmt.Println(fredPerson)

	// Print out the address of a Person object
	fmt.Println(&Person{name: "Ann", age: 40})

	// Print out the properties of a Person instance
	s := Person{name: "Sean", age: 50}
	fmt.Println(s.name, " ", s.age)
}
