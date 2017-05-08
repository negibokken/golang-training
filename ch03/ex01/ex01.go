package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-swidth: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, av := corner(float64(i+1), float64(j))
			bx, by, bv := corner(float64(i), float64(j))
			cx, cy, cv := corner(float64(i), float64(j+1))
			dx, dy, dv := corner(float64(i+1), float64(j+1))
			if !av || !bv || !cv || !dv {
				continue
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' />\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j float64) (float64, float64, bool) {
	valid := true
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	if math.IsNaN(sx) || math.IsInf(sx, 0) || math.IsNaN(sy) || math.IsInf(sy, 0) {
		valid = false
	}
	return sx, sy, valid
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
