package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

var outType string

func init() {
	flag.StringVar(&outType, "type", "jpeg", "convert to jpeg ,png or gif")
}

var out io.Writer = os.Stdout
var errout io.Writer = os.Stderr

func main() {
	flag.Parse()
	img, kind, err := image.Decode(os.Stdin)
	if err != nil {
		fmt.Fprintln(errout, "Input format =", kind)
	}
	fmt.Fprintln(errout, "Input format =", kind)
	if outType == kind {
		fmt.Fprintln(errout, "File type is already", kind)
		os.Exit(1)
	}
	handleType(outType, img, out)
}

func handleType(outType string, img image.Image, out io.Writer) (err error) {
	switch outType {
	case "jpeg", "jpg":
		err = toJPEG(img, out)
	case "png":
		err = toPNG(img, out)
	case "gif":
		err = toGIF(img, out)
	default:
		fmt.Fprintln(errout, "unknown type specified")
	}
	if err != nil {
		fmt.Fprintf(errout, "%v", err)
		return err
	}
	fmt.Fprintf(errout, "converted to %v\n", outType)
	return nil
}

var toJPEG = func(img image.Image, out io.Writer) error {
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}

var toPNG = func(img image.Image, out io.Writer) error {
	return png.Encode(out, img)
}

var toGIF = func(img image.Image, out io.Writer) error {
	return gif.Encode(out, img, &gif.Options{})
}
