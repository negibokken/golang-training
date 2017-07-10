package popcount

import "testing"

func TestPopCount(t *testing.T) {
	type args struct {
		x uint64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test popcount", args{10}, 2},
		{"test popcount", args{256}, 1},
		{"test popcount", args{255}, 8},
		{"test popcount", args{4095}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PopCount(tt.args.x); got != tt.want {
				t.Errorf("PopCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
