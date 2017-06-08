package main

import "testing"

func Test_containsAll(t *testing.T) {
	type args struct {
		x []string
		y []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test containsAll",
			args{
				[]string{""},
				[]string{""},
			},
			true,
		},
		{
			"test containsAll",
			args{
				[]string{"a", "b"},
				[]string{},
			},
			true,
		},
		{
			"test containsAll",
			args{
				[]string{"a", "b"},
				[]string{"c"},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsAll(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("containsAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasTarget(t *testing.T) {
	type args struct {
		attrStack []attr
		id        string
		value     string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test has Target",
			args{
				[]attr{
					attr{"id", "id-test"},
					attr{"class", "class-test"},
				},
				"id",
				"id-test",
			},
			true,
		},
		{
			"test has Target",
			args{
				[]attr{
					attr{"id", "id-test"},
					attr{"class", "class-test"},
				},
				"class",
				"id-test",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasTarget(tt.args.attrStack, tt.args.id, tt.args.value); got != tt.want {
				t.Errorf("hasTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
