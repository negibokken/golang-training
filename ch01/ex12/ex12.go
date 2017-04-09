package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
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
	var parameters = []string{"cycles", "size", "nframes", "delay"}
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
	lissajous(w, queries["cycles"], queries["size"], queries["nframes"], queries["delay"])
})

func lissajous(out io.Writer, _cycles int, _size int, _nframes int, _delay int) {
	var (
		cycles  = 5
		res     = 0.01
		size    = 100
		nframes = 5
		delay   = 8
	)
	if _cycles != 0 {
		cycles = _cycles
	}
	if _size != 0 {
		size = _size
	}
	if _nframes != 0 {
		nframes = _nframes
	}
	if _delay != 0 {
		delay = _delay
	}

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2.0*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), lineIndex)
			anim.Delay = append(anim.Delay, delay)
			anim.Image = append(anim.Image, img)
		}
		gif.EncodeAll(out, &anim)
	}
}
