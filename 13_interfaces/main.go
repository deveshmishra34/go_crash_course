package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	r float64
}

type Rectangle struct {
	h float64
	w float64
}

func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r Rectangle) area() float64 {
	return r.h * r.w
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{r: 5}
	rectangle := Rectangle{w: 10, h: 5}
	fmt.Printf("Circle area: %f \n", getArea(circle))
	fmt.Printf("Rectangle area: %f \n", getArea(rectangle))
}
