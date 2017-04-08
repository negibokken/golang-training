package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

// var palette = []color.Color{color.Black, color.RGBA{0x00, 0xff, 0x00, 0xff}}
var palette = []color.Color{}

const (
	bgIndex   = 0
	lineIndex = 1
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func mapToColor() []color.Color {
	var rgb [3]uint8
	var _palette []color.Color
	for i := 0; i < 256; i++ {
		origin := i
		for j := 0; j < 3; j++ {
			rgb[j] = uint8(origin * (j + 1) % 0xff)
			origin = origin % 0xff
		}
		col := color.RGBA{rgb[0], rgb[1], rgb[2], 0xff}
		_palette = append(_palette, col)
	}
	return _palette
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.01
		size    = 100
		nframes = 5
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		palette := mapToColor()
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			lineIndex := uint8(t/res*x*y) % 255
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), lineIndex)
			anim.Delay = append(anim.Delay, delay)
			anim.Image = append(anim.Image, img)
		}
		gif.EncodeAll(out, &anim)
	}
}
