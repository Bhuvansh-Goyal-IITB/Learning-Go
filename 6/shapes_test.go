package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10.0, 10.0}
	got := Perimeter(rect)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Height: 12.0, Width: 6.0}, want: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, want: 36},
	}

	for _, ts := range areaTests {
		t.Run(ts.name, func(t *testing.T) {
			got := ts.shape.Area()
			if got != ts.want {
				t.Errorf("%#v: got %g want %g", ts.shape, got, ts.want)
			}
		})
	}
}
