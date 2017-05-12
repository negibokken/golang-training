package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	superSampling(img)
	png.Encode(os.Stdout, img)
}

func averagePixels(img *image.RGBA, px int, py int) color.Color {
	var r, g, b, a uint32
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			rr, gg, bb, aa := img.At(px+j, py+i).RGBA()
			r += rr
			g += gg
			b += bb
			a += aa
		}
	}
	return color.RGBA{uint8(r / 4), uint8(g / 4), uint8(b / 4), uint8(a / 4)}
}

func superSampling(img *image.RGBA) {
	for py := 0; py < height-1; py++ {
		for px := 0; px < width-1; px++ {
			col := averagePixels(img, px, py)
			img.Set(px, py, col)
		}
	}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{250 - contrast*n*2, 240 - contrast*n, 230 - contrast*n*3, 255}
		}
	}
	return color.Black
}
