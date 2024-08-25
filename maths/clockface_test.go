package clockface_test

import (
	"math"
	"testing"
	"time"

	"github.com/hnucamendi/learn-go-with-tests/maths/clockface"
)

func testName(t time.Time) string {
	return t.Format("15:04:05")
}

func simpleTime(hour, minute, second int) time.Time {
	return time.Date(312, time.October, 28, hour, minute, second, 0, time.UTC)
}

func roughlyEqualFloat64(a, b float64) bool {
	equalityThreshold := 1e-7
	return math.Abs(a-b) < equalityThreshold
}

func roughlyEqualPoint(a, b clockface.Point) bool {
	return roughlyEqualFloat64(a.X, b.X) && roughlyEqualFloat64(a.Y, b.Y)
}

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{time: simpleTime(0, 0, 30), angle: math.Pi},
		{time: simpleTime(0, 0, 0), angle: 0},
		{time: simpleTime(0, 0, 45), angle: (math.Pi / 2) * 3},
		{time: simpleTime(0, 0, 7), angle: (math.Pi / 30) * 7},
	}

	for _, tt := range cases {
		t.Run(testName(tt.time), func(t *testing.T) {
			got := clockface.secondsInRadians(tt.time)
			if got != tt.angle {
				t.Fatalf("Wanted %v radians, but got %v", tt.angle, got)
			}
		})
	}
}

func TestSecondHandVector(t *testing.T) {
	cases := []struct {
		time  time.Time
		point clockface.Point
	}{
		{time: simpleTime(0, 0, 30), point: clockface.Point{X: 0, Y: -1}},
		{time: simpleTime(0, 0, 45), point: clockface.Point{X: -1, Y: 0}},
		{time: time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC), point: clockface.Point{}},
	}

	for _, tt := range cases {
		t.Run(testName(tt.time), func(t *testing.T) {
			got := clockface.secondHandPoint(tt.time)
			if !roughlyEqualPoint(tt.point, tt.point) {
				t.Fatalf("Wanted %v Point, but got %v", tt.point, got)
			}
		})
	}
}
