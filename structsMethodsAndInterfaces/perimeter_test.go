package structsmethodsandinterfaces

import "testing"

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "rectangle", shape: Rectangle{Width: 10, Height: 25}, want: 250.00},
		{name: "circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "triangle", shape: Triangle{Base: 32, Height: 23}, want: 368.00},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %g hasArea %g", tt.shape, got, tt.want)
			}
		})
	}
}

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "rectangle", shape: Rectangle{Width: 10, Height: 25}, want: 70},
		{name: "circle", shape: Circle{Radius: 25}, want: 157.07963267948966},
		{name: "triangle", shape: Triangle{A: 12, B: 23, Base: 32}, want: 67},
	}

	for _, tt := range perimeterTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Perimeter()
			if got != tt.want {
				t.Errorf("%#v got %g hasArea %g", tt.shape, got, tt.want)
			}
		})
	}
}
