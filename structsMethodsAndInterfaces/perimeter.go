package structsmethodsandinterfaces

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Triangle struct {
	A      float64
	B      float64
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}

func (t Triangle) Perimeter() float64 {
	return t.A + t.B + t.Base
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Perimeter() float64 {
	return (c.Radius * math.Pi) * 2
}

func (c Circle) Area() float64 {
	return math.Pow(c.Radius, 2) * math.Pi
}
