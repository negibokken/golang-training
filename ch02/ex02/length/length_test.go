package length

import "testing"

// Test for Celcius
func TestMToI(t *testing.T) {
	tests := []struct {
		m        Meter
		expected Inch
	}{
		{0.508, 20},
		{3.048, 120},
	}

	for _, test := range tests {
		if MToI(test.m) != test.expected {
			t.Errorf("Not much %v and %v", test.m, MToI(test.m))
		}
	}
}

func TestMToF(t *testing.T) {
	tests := []struct {
		c        Meter
		expected Feet
	}{
		{3.048, 10},
		{16.764, 55},
	}

	for _, test := range tests {
		if MToF(test.c) != test.expected {
			t.Errorf("Not much %v and %v", test.c, MToF(test.c))
		}
	}
}

func TestMString(t *testing.T) {
	tests := []struct {
		m        Meter
		expected string
	}{
		{0, "0m"},
		{10, "10m"},
		{20, "20m"},
	}

	for _, test := range tests {
		if test.m.String() != test.expected {
			t.Errorf("Not much")
		}
	}
}

// Test for Inch
func TestIToM(t *testing.T) {
	tests := []struct {
		i        Inch
		expected Meter
	}{
		{1, 0.0254},
		{100, 2.54},
		{500, 12.7},
	}

	for _, test := range tests {
		if IToM(test.i) != test.expected {
			t.Errorf("Not much %v and %v", test.i, IToM(test.i))
		}
	}
}

func TestIToF(t *testing.T) {
	tests := []struct {
		f        Inch
		expected Feet
	}{
		{60, 5},
		{240, 20},
		{360, 30},
	}

	for _, test := range tests {
		if IToF(test.f) != test.expected {
			t.Errorf("Not much %v and %v", test.f, IToF(test.f))
		}
	}
}

func TestIString(t *testing.T) {
	tests := []struct {
		i        Inch
		expected string
	}{
		{0, "0in"},
		{10, "10in"},
		{20, "20in"},
	}

	for _, test := range tests {
		if test.i.String() != test.expected {
			t.Errorf("Not much")
		}
	}

}

// Test for Feet
func TestFToM(t *testing.T) {
	tests := []struct {
		f        Feet
		expected Meter
	}{
		{10, 3.048},
		{30, 9.144},
		{40, 12.192},
	}

	for _, test := range tests {
		if FToM(test.f) != test.expected {
			t.Errorf("Not much %v and %v", test.f, FToM(test.f))
		}
	}
}

func TestFToI(t *testing.T) {
	tests := []struct {
		k        Feet
		expected Inch
	}{
		{5, 60},
		{20, 240},
		{40, 480},
	}

	for _, test := range tests {
		if FToI(test.k) != test.expected {
			t.Errorf("Not much %v and %v", test.k, FToI(test.k))
		}
	}
}

func TestFString(t *testing.T) {
	tests := []struct {
		f        Feet
		expected string
	}{
		{0, "0ft"},
		{10, "10ft"},
		{20, "20ft"},
	}

	for _, test := range tests {
		if test.f.String() != test.expected {
			t.Errorf("Not much")
		}
	}
}
