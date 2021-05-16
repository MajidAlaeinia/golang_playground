package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width, Height float64
}

func (c Circle) area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func (r Rectangle) area() float64 {
	return r.Width * r.Height
}

func getArea(sh Shape) float64 {
	return sh.area()
}

type Animal interface {
	seda() string
}

type Cat struct {
	Name string
}

type Dog struct {
	Name string
}

func (c Cat) seda() string {
	return "Meow"
}

func (d Dog) seda() string {
	return "Hop"
}

func getSeda(a Animal) string {
	return a.seda()
}

func main() {
	c := Cat{"Sirus"}
	d := Dog{"Virus"}

	fmt.Println(getSeda(c))
	fmt.Println(getSeda(d))

	ci := Circle{5}
	re := Rectangle{5, 4}

	fmt.Println(getArea(ci))
	fmt.Println(getArea(re))
}
