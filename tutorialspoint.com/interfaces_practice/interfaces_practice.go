package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

func (circle Circle) area() float64 {
	return math.Pi * math.Pow(circle.radius, 2)
}

func (rectangle Rectangle) area() float64 {
	return rectangle.width * rectangle.height
}

func getArea(shape Shape) float64 {
	return shape.area()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{width: 10, height: 5}

	fmt.Printf("Circle Area: %0.2f\n", getArea(circle))
	fmt.Printf("Rectangle Area: %0.2f\n", getArea(rectangle))
}