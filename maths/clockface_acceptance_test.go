package clockface_test

import (
	"bytes"
	"encoding/xml"
	"testing"
	"time"

	clockface "github.com/hnucamendi/learn-go-with-tests/maths"
)

func containsLine(line Line, lines []Line) bool {
	for _, l := range lines {
		if line == l {
			return true
		}
	}
	return false
}

func TestSVGWriterAtMidnight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)
	b := bytes.Buffer{}

	clockface.SVGWriter(&b, tm)

	svg := &SVG{}

	xml.Unmarshal(b.Bytes(), &svg)

	want := Line{X1: 150, X2: 150, Y1: 150, Y2: 60}

	for _, line := range svg.Line {
		if line == want {
			return
		}
	}

	t.Errorf("Expected to find the second hand line %+v, in the SVG lines %+v", want, svg.Line)
}

func TestSVGWriterSecondHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{time: simpleTime(0, 0, 0), line: Line{X1: 150, X2: 150, Y1: 150, Y2: 60}},
		{time: simpleTime(0, 0, 30), line: Line{X1: 150, X2: 150, Y1: 150, Y2: 240}},
	}

	for _, tt := range cases {
		t.Run(testName(tt.time), func(t *testing.T) {
			b := bytes.Buffer{}
			clockface.SVGWriter(&b, tt.time)

			svg := &SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(tt.line, svg.Line) {
				t.Errorf("Expected to find the second hand line %+v, in the SVG lines %+v", tt.line, svg.Line)
			}
		})
	}
}

func TestSVGWriterMinutedHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{
			simpleTime(0, 0, 0),
			Line{150, 150, 150, 70},
		},
	}

	for _, tt := range cases {
		t.Run(testName(tt.time), func(t *testing.T) {
			b := bytes.Buffer{}
			clockface.SVGWriter(&b, tt.time)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(tt.line, svg.Line) {
				t.Errorf("Expected to find the minute hand line %+v, in the SVG lines %+v", tt.line, svg.Line)
			}
		})
	}
}

func TestSVGWriterHourHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{
			simpleTime(6, 0, 0),
			Line{150, 150, 150, 200},
		},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			b := bytes.Buffer{}
			clockface.SVGWriter(&b, c.time)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(c.line, svg.Line) {
				t.Errorf("Expected to find the hour hand line %+v, in the SVG lines %+v", c.line, svg.Line)
			}
		})
	}
}
