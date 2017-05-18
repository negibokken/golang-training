package main

import (
	"testing"
)

func Test_printUsageAndExit(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"name"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printUsageAndExit()
		})
	}
}
