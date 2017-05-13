package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.Black, color.RGBA{0x00, 0xff, 0x00, 0xff}}

const (
	lineIndex = 1
)

const hostName = "localhost:8000"

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(hostName, nil))
}

type Query struct {
	x int
	y int
	m float64
}

var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var queries Query
	for key, value := range r.URL.Query() {
		if key == "x" {
			num, err := strconv.Atoi(value[0])
			if err != nil {
				fmt.Fprintf(w, "%v", err)
			}
			queries.x = num
		} else if key == "y" {
			num, err := strconv.Atoi(value[0])
			if err != nil {
				fmt.Fprintf(w, "%v", err)
			}
			queries.y = num
		} else if key == "m" {
			num, err := strconv.ParseFloat(value[0], 64)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
			}
			queries.m = num
		}
	}
	w.Header().Set("Content-Type", "image/png")
	fractale(w, queries.x, queries.y, queries.m)
})

func fractale(w io.Writer, x, y int, m float64) {
	xmin, ymin, xmax, ymax := float64(x-2)/m, float64(y-2)/m, float64(x+2)/m, float64(x+2)/m
	width, height := 1024, 1024
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(w, img)
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
