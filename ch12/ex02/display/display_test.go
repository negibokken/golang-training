package display

import (
	"testing"
)

func TestDisplay(t *testing.T) {
	type args struct {
		name string
		x    interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{"int", args{"int", 1}},
		{"string", args{"string", "string"}},
		{"map string array", args{"map", map[string]string{"map": "key"}}},
		{"name", args{"map", map[[2]string]string{[2]string{"map1", "map2"}: "key"}}},
		{"name", args{"map", map[struct{ int }]string{{1}: "key"}}},
		{"name", args{"map", map[struct{ string }]string{{"string"}: "key"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Display(tt.args.name, tt.args.x)
		})
	}
}
