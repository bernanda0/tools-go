package abstract

import (
	"fmt"
	"math"
)

// for the projec
type BaseResponse struct {
	Err      error
	PanicMsg interface{}
}

type Square struct {
	height, width float64
}

type Circle struct {
	radius float64
}

type Shape interface {
	area() float64
	perimeter() float64
}

func (square Square) area() float64 {
	return square.height * square.width
}

func (square Square) perimeter() float64 {
	return 2 * (square.height + square.width)
}

func (circle Circle) area() float64 {
	return math.Pi * math.Pow(circle.radius, 2)
}

// Function that has parameter of the interface
func getArea(shape Shape) float64 {
	return shape.area()
}

func getPerimeter(shape Shape) float64 {
	return shape.perimeter()
}

func AbstractRun() {
	s := Square{12, 14}
	c := Circle{7}

	fmt.Println(getArea(s))
	fmt.Println(getPerimeter(s))

	fmt.Println(s.area())
	fmt.Println(s.perimeter())

	// fmt.Println(getArea(c))
	fmt.Println(c.area())
}
