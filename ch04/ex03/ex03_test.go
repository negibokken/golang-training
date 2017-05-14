package main

import "testing"

func TestReverse(t *testing.T) {
	type args struct {
		num *[size]int
	}
	tests := []struct {
		name   string
		args   args
		expect [size]int
	}{
		{
			"逆順",
			args{&[size]int{3, 4, 5, 6, 7}},
			[size]int{7, 6, 5, 4, 3},
		},
		{
			"逆順",
			args{&[size]int{13, 14, 15, 16, 17}},
			[size]int{17, 16, 15, 14, 13},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverse(tt.args.num)
			if *tt.args.num != tt.expect {
				t.Errorf("got: %v expect: %v", tt.args.num, tt.expect)
			}
		})
	}
}
