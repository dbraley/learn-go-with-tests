package geometry

type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter of a rectangle
func Perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}

// Area of a rectangle
func Area(r Rectangle) float64 {
	return r.Width * r.Height
}
