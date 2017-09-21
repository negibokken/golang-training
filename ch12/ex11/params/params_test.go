package params

import (
	"testing"
)

func TestPack(t *testing.T) {
	type Profile struct {
		Name string `http:"n"`
		Age  int    `http:"a"`
	}
	tests := []struct {
		p      Profile
		expect string
	}{
		{Profile{"Alice", 20}, "a=20&n=Alice"},
		{Profile{"Bob", 21}, "a=21&n=Bob"},
		{Profile{"Charlie", 22}, "a=22&n=Charlie"},
	}
	for _, test := range tests {
		u, err := Pack(&test.p)
		if err != nil {
			t.Errorf("Pack(%#v): %s", test.p, err)
		}
		got := u.RawQuery
		if got != test.expect {
			t.Errorf("Pack(%#v): got %q, want %q", test.p, got, test.expect)
		}
	}
}
