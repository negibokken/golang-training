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

var out = os.Stdout

func main() {
	flag.Parse()
	img, kind, err := image.Decode(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Input format =", kind)
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	if outType == kind {
		fmt.Fprintln(os.Stderr, "File type is already", kind)
		os.Exit(1)
	}
	handleType(outType, img, out)
}

func handleType(outType string, img image.Image, out io.Writer) (err error) {
	switch outType {
	case "jpeg":
		err = toJPEG(img, out)
	case "png":
		err = toPNG(img, out)
	case "gif":
		err = toGIF(img, out)
	default:
		fmt.Fprintln(os.Stderr, "unknown type specified")
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return err
	}
	fmt.Fprintf(os.Stderr, "conver ted to %v\n", outType)
	return nil
}

func toJPEG(img image.Image, out io.Writer) error {
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}

func toPNG(img image.Image, out io.Writer) error {
	return png.Encode(out, img)
}

func toGIF(img image.Image, out io.Writer) error {
	return gif.Encode(out, img, &gif.Options{})
}
