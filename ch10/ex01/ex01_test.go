package main

import (
	"image"
	"io"
	"os"
	"testing"
)

var temp = func(img image.Image, out io.Writer) error {
	return nil
}

func Test_handleType(t *testing.T) {
	f, err := os.Open("gopher.gif")
	if err != nil {
		t.Errorf("cannot open file")
		return
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		t.Errorf("cannot open file")
		return
	}
	type args struct {
		outType string
		img     image.Image
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
		wantErr bool
	}{
		{"outType", args{"jpeg", img}, "converted to jpeg", false},
	}

	// Make stub
	originJ := toJPEG
	toJPEG = temp
	defer func() { toJPEG = originJ }()
	originP := toJPEG
	toPNG = temp
	defer func() { toPNG = originP }()
	originG := toGIF
	toGIF = temp
	defer func() { toGIF = originG }()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errout = (*os.File)(nil)
			if err := handleType(tt.args.outType, tt.args.img, out); (err != nil) != tt.wantErr {
				t.Errorf("handleType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if gotOut := errout.String(); gotOut != tt.wantOut {
			// t.Errorf("handleType() = %v, want %v", gotOut, tt.wantOut)
			// }
		})
	}
}
