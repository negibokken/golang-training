package main

import (
	"fmt"
	"image/color"
	"io"
	"log"
	"math"
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

var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var parameters = []string{"width", "height", "cells"}
	queries := map[string]int{}
	for _, param := range parameters {
		queries[param] = 0
	}
	for key, value := range r.URL.Query() {
		num, err := strconv.Atoi(value[0])
		if err != nil {
			fmt.Fprintf(w, "%v", err)
		}
		queries[key] = num
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	svgGenerator(w, queries["width"], queries["height"], queries["cells"])
})

const angle = math.Pi / 6

var width int
var height int
var xyscale int
var zscale int
var xyrange float64
var cells int
var sin30, cos30 float64 = math.Sin(angle), math.Cos(angle)

func svgGenerator(out io.Writer, _width int, _height int, _cells int) {
	height = _height
	width = _width
	height = _height
	if width == 0 {
		width = 600
	}
	if height == 0 {
		height = 320
	}
	if xyrange == 0.0 {
		xyrange = 30.0
	}
	if cells == 0 {
		cells = 100
	}
	xyscale = width / 2 / int(xyrange)
	zscale = int(float64(height) * 0.4)

	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; stroke-swidth: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az, av := corner(float64(i+1), float64(j))
			bx, by, bz, bv := corner(float64(i), float64(j))
			cx, cy, cz, cv := corner(float64(i), float64(j+1))
			dx, dy, dz, dv := corner(float64(i+1), float64(j+1))
			if !av || !bv || !cv || !dv {
				continue
			}
			zave := (az + bz + cz + dz) / 4
			color := calculateColor(float64(zave))

			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='#%v' />\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Fprintf(out, "</svg>")
}

func calculateColor(z float64) string {
	const maxHeight = 0.33
	color := uint((0xff0000 * float64((z / 0.33))) + 0x0000ff)
	return fmt.Sprintf("%06x", color)
}

func corner(i, j float64) (float64, float64, float64, bool) {
	valid := true
	x := xyrange * (float64(i)/float64(cells) - 0.5)
	y := xyrange * (float64(j)/float64(cells) - 0.5)

	z := f(x, y)

	sx := float64(width)/2 + (x-y)*cos30*float64(xyscale)
	sy := float64(height)/2 + (x+y)*sin30*float64(xyscale) - z*float64(zscale)
	if math.IsNaN(sx) || math.IsInf(sx, 0) || math.IsNaN(sy) || math.IsInf(sy, 0) {
		valid = false
	}
	return sx, sy, z, valid
}

func f(x, y float64) float64 {
	return math.Pow(2, math.Sin(y)) * math.Pow(2, math.Sin(x)) / 12
}
