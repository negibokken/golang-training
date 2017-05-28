package main

import "testing"

func Test_foo(t *testing.T) {
	tests := []struct {
		name       string
		wantResult string
	}{
		{
			"test foo",
			"bar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := foo(); gotResult != tt.wantResult {
				t.Errorf("foo() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
