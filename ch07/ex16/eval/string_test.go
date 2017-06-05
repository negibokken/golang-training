package eval

import (
	"testing"
)

func TestExpr(t *testing.T) {
	tests := []struct {
		name string
		expr string
		want string
	}{
		{"expr test", "1+1", "1+1"},
		{"expr test", "pow(2, 3)", "pow(2, 3)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ex, err := Parse(tt.expr)
			if err != nil {
				t.Errorf("%v", err)
			}
			ex2, err := Parse(ex.String())
			if err != nil {
				t.Errorf("%v", err)
			}
			if ex2.String() != tt.want {
				t.Errorf("got: %v, want: %v", ex2.String(), tt.want)
			}
		})
	}
}
