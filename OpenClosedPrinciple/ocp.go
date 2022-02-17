package main

import (
	"fmt"
	"math"
)

// open closed principle
// open for extension, closed for modification

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t *Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

//create a struct
type ShapeFactory struct{}

//create a method that takes a shape interface and returns the area
func (sf *ShapeFactory) GetArea(s Shape) float64 {
	fmt.Println(s.Area())
	return s.Area()
}

func main() {

	shape1 := ShapeFactory{}
	rect := Rectangle{
		Width:  10,
		Height: 20,
	}
	shape1.GetArea(&rect)
}
