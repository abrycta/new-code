package main

import (
	"fmt"
	"math"
)

/*
interface for geometric shapes
*/
type Geometry interface {
	area() float64
	perim() float64
}

/*
implementing the interface on geometric shapes Circle and Rectangle
*/
type Rect struct {
	width, height float64
}
type Circle struct {
	radius float64
}

/*
Implementation for Rectangle
*/
func (r Rect) area() float64 {
	return r.width * r.height
}
func (r Rect) perim() float64 {
	return 2*r.width + 2*r.height
}

/*
Implementation for Circle
*/
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

/*
calling methods in named interface
*/
func measure(g Geometry) {
	fmt.Println(g)         // prints the inserted values
	fmt.Println(g.area())  // prints the area of the geometric shape
	fmt.Println(g.perim()) // prints the perimeter of the geometric shape
}

/*
Main method where Circle and Rectangle implements the Geometry interface as arguments to function measure.
*/
func main() {
	r := Rect{width: 3, height: 4} // input height and width of a Rectangle
	c := Circle{radius: 5}         // input radius of a Circle

	measure(r) // call measure function
	measure(c)
}
