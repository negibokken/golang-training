package main

import (
	"math"
	"testing"
)

func TestCorner(t *testing.T) {
	// +inf, -inf and NaN should be checked
	var normalTests = []struct {
		i float64
		j float64
	}{
		{0, 0},
		{1, 1},
		{-1, -1},
	}

	for _, test := range normalTests {
		_, _, v := corner(test.i, test.j)
		if v == false {
			t.Errorf("Unexpected false. corner(%v,%v) should not return false", test.i, test.j)
		}
	}

	var invalidTests = []struct {
		i float64
		j float64
	}{
		{math.Inf(1), 0},
		{math.Inf(-1), 0},
		{math.NaN(), 0},
	}

	for _, test := range invalidTests {
		_, _, v := corner(test.i, test.j)
		if v == true {
			t.Errorf("Unexpected false. corner(%v,%v) should not return false", test.i, test.j)
		}
	}
}
