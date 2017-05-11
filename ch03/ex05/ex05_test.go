package main

import (
	"image/color"
	"testing"
)

func TestMandelbrot(t *testing.T) {
	var tests = []struct {
		z        complex128
		expected color.Color
	}{
		{1 + 1i, color.RGBA{220, 225, 185, 255}},
		{2 + 1i, color.RGBA{250, 240, 230, 255}},
		{0.002 + 0.00002i, color.Black},
	}

	for _, test := range tests {
		if actual := mandelbrot(test.z); actual != test.expected {
			t.Errorf("mandelbrot(%v) should be %v but %v", test.z, actual, test.expected)
		}
	}
}
