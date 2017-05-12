package main

import (
	"image/color"
	"reflect"
	"testing"
)

func TestNewton(t *testing.T) {
	type args struct {
		z complex128
	}
	tests := []struct {
		args args
		want color.Color
	}{
		{args{complex(1, 2)}, color.Gray{234}},
		{args{complex(0.1, 0.3)}, color.Gray{210}},
	}
	for _, tt := range tests {
		if got := newton(tt.args.z); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("newton() = %v, want %v", got, tt.want)
		}
	}
}
