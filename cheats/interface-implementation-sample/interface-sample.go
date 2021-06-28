package interfaceimplementationsample

import "fmt"

// Interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Struct
type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

// Implementation interface of Shape is a contract,
// Struct Rectangle should have method Area and Perimeter
// Then it struct can be able to assign to variable shape
func Implementation() error {
	var shape Shape
	shape = Rectangle{
		width:  4.0,
		height: 5.0,
	}
	fmt.Println("area of rectangle is ", shape.Area())
	fmt.Println("perimeter of rectangle is ", shape.Perimeter())
	return nil
}
