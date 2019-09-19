package geometry

import "math"

// Shape contains floating point geometric caluclations
type Shape interface {
	Area() float64
}

// Rectangle has a width and height
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle has a radius
type Circle struct {
	Radius float64
}

// Triangle has a base and height
type Triangle struct {
	Base   float64
	Height float64
}

// Perimeter of a rectangle
func Perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}

// Area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area of a Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Area of a Triangle
func (t Triangle) Area() float64 {
	return t.Base * t.Height * 0.5
}
