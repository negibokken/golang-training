package main

import (
	"image"
	"image/color"
	"reflect"
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

func TestAveragePixels(t *testing.T) {
	width := 1024
	height := 1024
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	var tests = []struct {
		pixels   []color.RGBA
		expected color.RGBA
	}{
		{
			[]color.RGBA{
				color.RGBA{0, 0, 0, 255},
				color.RGBA{0, 0, 0, 255},
				color.RGBA{0, 0, 0, 255},
				color.RGBA{0, 0, 0, 255},
			},
			color.RGBA{0, 0, 0, 255},
		},
		{
			[]color.RGBA{
				color.RGBA{1, 5, 9, 255},
				color.RGBA{2, 6, 10, 255},
				color.RGBA{3, 7, 11, 255},
				color.RGBA{6, 2, 14, 255},
			},
			color.RGBA{3, 5, 11, 255},
		},
	}

	for _, test := range tests {
		assignPixel(img, test.pixels)
		if pixel := averagePixels(img, 0, 0); pixel != test.expected {
			t.Errorf("averagePixes should be %v but %v", test.expected, pixel)
		}
	}
}

func assignPixel(img *image.RGBA, pixels []color.RGBA) {
	cnt := 0
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			img.Set(i, j, pixels[cnt])
			cnt++
		}
	}
}

func Test_averagePixels(t *testing.T) {
	type args struct {
		img *image.RGBA
		px  int
		py  int
	}
	tests := []struct {
		name string
		args args
		want color.Color
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := averagePixels(tt.args.img, tt.args.px, tt.args.py); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("averagePixels() = %v, want %v", got, tt.want)
			}
		})
	}
}
