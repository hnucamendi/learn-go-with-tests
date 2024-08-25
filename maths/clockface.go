package clockface

import (
	"math"
	"time"
)

const (
	secondHandLength float64 = 90
	clockCentreX     float64 = 150
	clockCentreY     float64 = 150
)

// A Point represents a two dimensional Cartesian coordinate
type Point struct {
	X float64
	Y float64
}

// SecondHand is the unit vector of the second hand of an analogue clock at time `t`
// represented as a Point.
func SecondHand(t time.Time) Point {
	p := secondHandPoint(t)
	p = Point{X: p.X * secondHandLength, Y: p.Y * secondHandLength} // scale
	p = Point{X: p.X, Y: -p.Y}                                      // flip
	p = Point{X: p.X + clockCentreX, Y: p.Y + clockCentreY}         // translate
	return p
}

func secondsInRadians(t time.Time) float64 {
	return (math.Pi / (30 / float64(t.Second())))
}

func secondHandPoint(t time.Time) Point {
	angle := secondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{X: x, Y: y}
}
