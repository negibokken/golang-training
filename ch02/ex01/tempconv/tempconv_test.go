package tempconv

import "testing"

// Test for Celcius
func TestCToF(t *testing.T) {
	tests := []struct {
		c        Celsius
		expected Fahrenheit
	}{
		{0, 32},
		{10, 50},
		{20, 68},
	}

	for _, test := range tests {
		if CToF(test.c) != test.expected {
			t.Errorf("Not much %v and %v", test.c, CToF(test.c))
		}
	}
}

func TestCToK(t *testing.T) {
	tests := []struct {
		c        Celsius
		expected Kelvin
	}{
		{0, 273.15},
		{10, 283.15},
		{20, 293.15},
	}

	for _, test := range tests {
		if CToK(test.c) != test.expected {
			t.Errorf("Not much %v and %v", test.c, CToK(test.c))
		}
	}
}

func TestCString(t *testing.T) {
	tests := []struct {
		c        Celsius
		expected string
	}{
		{0, "0°C"},
		{10, "10°C"},
		{20, "20°C"},
	}

	for _, test := range tests {
		if test.c.String() != test.expected {
			t.Errorf("Not much")
		}
	}
}

// Test for Fahrenheit
func TestFToC(t *testing.T) {
	tests := []struct {
		f        Fahrenheit
		expected Celsius
	}{
		{5, -15},
		{14, -10},
		{23, -5},
	}

	for _, test := range tests {
		if FToC(test.f) != test.expected {
			t.Errorf("Not much %v and %v", test.f, FToC(test.f))
		}
	}
}

func TestFToK(t *testing.T) {
	tests := []struct {
		f        Fahrenheit
		expected Kelvin
	}{
		{5, 258.15},
		{14, 263.15},
		{23, 268.15},
	}

	for _, test := range tests {
		if FToK(test.f) != test.expected {
			t.Errorf("Not much %v and %v", test.f, FToK(test.f))
		}
	}
}

func TestFString(t *testing.T) {
	tests := []struct {
		f        Fahrenheit
		expected string
	}{
		{0, "0°F"},
		{10, "10°F"},
		{20, "20°F"},
	}

	for _, test := range tests {
		if test.f.String() != test.expected {
			t.Errorf("Not much")
		}
	}

}

// Test for Kelvin
func TestKToC(t *testing.T) {
	tests := []struct {
		k        Kelvin
		expected Celsius
	}{
		{10, -263.15},
		{0.15, -273},
		{0.75, -272.4},
	}

	for _, test := range tests {
		if KToC(test.k) != test.expected {
			t.Errorf("Not much %v and %v", test.k, KToC(test.k))
		}
	}
}

func TestKToF(t *testing.T) {
	tests := []struct {
		k        Kelvin
		expected Fahrenheit
	}{
		{273.15, 32},
		{283.15, 50},
		{339.9, 152.15},
	}

	for _, test := range tests {
		if KToF(test.k) != test.expected {
			t.Errorf("Not much %v and %v", test.k, KToF(test.k))
		}
	}
}

func TestKString(t *testing.T) {
	tests := []struct {
		k        Kelvin
		expected string
	}{
		{0, "0K"},
		{10, "10K"},
		{20, "20K"},
	}

	for _, test := range tests {
		if test.k.String() != test.expected {
			t.Errorf("Not much")
		}
	}
}
