package main

import (
	"reflect"
	"testing"
)

func TestDeleteRedundantSpace(t *testing.T) {
	type args struct {
		str []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"前方にスペース", args{[]byte("  Hello, world")}, []byte(" Hello, world")},
		{"中央にスペース", args{[]byte("Hello,  world")}, []byte("Hello, world")},
		{"後方にスペース", args{[]byte("Hello, world  ")}, []byte("Hello, world ")},
		{"前方と中央にスペース", args{[]byte("  Hello,  world")}, []byte(" Hello, world")},
		{"中央と後方にスペース", args{[]byte("Hello,  world  ")}, []byte("Hello, world ")},
		{"前方と中央と後方にスペース", args{[]byte("  Hello,  world  ")}, []byte(" Hello, world ")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteRedundantSpace(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteRedundantSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
