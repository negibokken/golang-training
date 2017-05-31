package main

import "testing"

func Test_tree_String(t *testing.T) {

	tests := []struct {
		name    string
		value   int
		t2Value int
		t3Value int
		want    string
	}{
		{"test tree string", 1, 2, 3, "213"},
		{"test tree string", 2, 3, 5, "325"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t2 := tree{tt.t2Value, nil, nil}
			t3 := tree{tt.t3Value, nil, nil}
			t1 := tree{tt.value, &t2, &t3}
			if got := t1.String(); got != tt.want {
				t.Errorf("tree.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
