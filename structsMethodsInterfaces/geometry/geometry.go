package geometry

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Length  float64
	Breadth float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Breadth
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * math.Pow(c.Radius, 2.0)
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Length + rectangle.Breadth)
}

func Area() float64 {
	return 0
}
