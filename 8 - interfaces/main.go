package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

// type SomeOtherType struct {
// 	name string
// }

type Rect struct {
	width  float64
	height float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

var s Shape

func main() {
	s = Rect{width: 5.0, height: 4.0}
	r := Rect{5.0, 4.0} // Don't have to use names

	fmt.Printf("type of s is %T\n", s)
	fmt.Printf("value of s is %v\n", s)
	fmt.Println("area of rectangle s", s.Area())
	fmt.Println("s == r is", s == r)

	//someTime := SomeOtherType{"eli"}
	// printArea(someTime)
	// cannot use someTime (variable of type SomeOtherType) as Shape value
	// in argument to printArea: missing method Areacompiler

	printArea(r)

}

func printArea(s Shape) {
	fmt.Printf("Area of shape is: %v", s.Area())
}
