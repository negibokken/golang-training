package main

import "testing"

func Test_max(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			"test max",
			args{[]int{1, 2, 3, 4, 5}},
			5,
			false,
		},
		{
			"test max",
			args{[]int{-1, -2, -3, -4, -5}},
			-1,
			false,
		},
		{
			"test max",
			args{[]int{}},
			0,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := max(tt.args.nums...)
			if (err != nil) != tt.wantErr {
				t.Errorf("max() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_min(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			"test min",
			args{[]int{1, 2, 3, 4, 5}},
			1,
			false,
		},
		{
			"test min",
			args{[]int{-1, -2, -3, -4, -5}},
			-5,
			false,
		},
		{
			"test min",
			args{[]int{}},
			0,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := min(tt.args.nums...)
			if (err != nil) != tt.wantErr {
				t.Errorf("min() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_max2(t *testing.T) {
	type args struct {
		n    int
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test max",
			args{0, []int{1, 2, 3, 4, 5}},
			5,
		},
		{
			"test max",
			args{0, []int{-1, -2, -3, -4, -5}},
			0,
		},
		{
			"test max",
			args{0, []int{}},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := max2(tt.args.n, tt.args.nums...); got != tt.want {
				t.Errorf("max2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_min2(t *testing.T) {
	type args struct {
		n    int
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test max",
			args{0, []int{1, 2, 3, 4, 5}},
			0,
		},
		{
			"test max",
			args{0, []int{-1, -2, -3, -4, -5}},
			-5,
		},
		{
			"test max",
			args{0, []int{}},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := min2(tt.args.n, tt.args.nums...); got != tt.want {
				t.Errorf("min2() = %v, want %v", got, tt.want)
			}
		})
	}
}
