package main

import "testing"

func TestRotate(t *testing.T) {
	type args struct {
		num *[size]int
		d   int
	}
	tests := []struct {
		name     string
		args     args
		expected [size]int
	}{
		{"Rotate", args{&[size]int{1, 2, 3, 4, 5}, 2}, [size]int{3, 4, 5, 1, 2}},
		{"Rotate", args{&[size]int{1, 2, 3, 4, 5}, 3}, [size]int{4, 5, 1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.num, tt.args.d)
			if *tt.args.num != tt.expected {
				t.Errorf("got: %v expected: %v", *tt.args.num, tt.expected)
			}
		})
	}
}
