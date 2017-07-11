package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/cmplx"
	"os"
	"sync"
)

var out io.Writer = os.Stdout

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	var wg sync.WaitGroup
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			wg.Add(1)
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			go func(ppx, ppy int, z complex128) {
				defer wg.Done()
				img.Set(ppx, ppy, mandelbrot(z))
			}(px, py, z)
		}
	}

	go func() {
		wg.Wait()
	}()

	png.Encode(out, img)
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
