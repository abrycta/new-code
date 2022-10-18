package main

import "fmt"

// demonstration of receiver functions,
// receiver functions = methods
// similar to instance methods:
// CarInstance.startEngine()
// We use 'rect' here because 'Rect' has already been
// declared in the same package
type rect struct {
	width, height int
}

// method that has a receiver type of *rect
func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	//call 2 methods defined for the struct
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	//use pointer receiver to avoid copying values on method calls
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
