package main

import (
	"testing"
)

func (s *IntSet) addElements(elements []int) {
	for _, ele := range elements {
		s.Add(ele)
	}
}

func TestIntSet_Has(t *testing.T) {
	type fields struct {
		words []int
	}
	type args struct {
		x int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"Has test", fields{[]int{1, 2, 3}}, args{1}, true},
		{"Has test", fields{[]int{1, 2, 3}}, args{2}, true},
		{"Has test", fields{[]int{1, 2, 3}}, args{3}, true},
		{"Has test", fields{[]int{1, 2, 3}}, args{4}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &IntSet{}

			for _, ele := range tt.fields.words {
				s.Add(ele)
			}
			if got := s.Has(tt.args.x); got != tt.want {
				t.Errorf("IntSet.Has() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntSet_Add(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name   string
		fields []int
		args   args
		want   []int
	}{
		{"Add test", []int{1}, args{2}, []int{1, 2}},
		{"Add test", []int{1, 2, 3}, args{2}, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &IntSet{}
			for _, ele := range tt.fields {
				s.Add(ele)
			}
			s.Add(tt.args.x)
			for _, w := range tt.want {
				if !s.Has(w) {
					t.Errorf("s doesn't have %v", w)
				}
			}
		})
	}
}

func TestIntSet_Len(t *testing.T) {
	tests := []struct {
		name   string
		fields []int
		want   int
	}{
		{"Len test", []int{1, 2, 3}, 3},
		{"Len test", []int{1, 2, 3, 4}, 4},
		{"Len test", []int{}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &IntSet{}
			for _, ele := range tt.fields {
				s.Add(ele)
			}
			if got := s.Len(); got != tt.want {
				t.Errorf("IntSet.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntSet_Remove(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name   string
		fields []int
		args   args
		want   []int
		dwant  []int
	}{
		{"Remove test", []int{1, 2, 3}, args{3}, []int{1, 2}, []int{3}},
		{"Remove test", []int{1, 2, 3}, args{2}, []int{1, 3}, []int{2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &IntSet{}
			for _, ele := range tt.fields {
				s.Add(ele)
			}
			s.Remove(tt.args.x)
			for _, w := range tt.want {
				if !s.Has(w) {
					t.Errorf("s doesn't have %v", w)
				}
			}
			for _, w := range tt.dwant {
				if s.Has(w) {
					t.Errorf("s have removed %v", w)
				}
			}
		})
	}
}

func TestIntSet_Clear(t *testing.T) {
	tests := []struct {
		name   string
		fields []int
	}{
		{"Clear test", []int{1, 2, 3}},
		{"Clear test", []int{1, 2, 3, 4, 5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &IntSet{}
			for _, ele := range tt.fields {
				s.Add(ele)
			}
			s.Clear()
			if s.Len() != 0 {
				t.Errorf("Clear didn't work")
			}
		})
	}
}

func TestIntSet_Copy(t *testing.T) {
	tests := []struct {
		name   string
		fields []int
		want   []int
	}{

		{"Copy test", []int{1, 2, 3}, []int{1, 2, 3}},
		{"Copy test", []int{}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &IntSet{}
			for _, ele := range tt.fields {
				s.Add(ele)
			}
			if s.Len() != len(tt.want) {
				t.Errorf("got len: %v, want len %v", s.Len(), len(tt.want))
			}
			for _, w := range tt.want {
				if !s.Has(w) {
					t.Errorf("s should have %v", w)
				}
			}
		})
	}
}

func TestIntSet_UnionWith(t *testing.T) {
	tests := []struct {
		name    string
		fields  []int
		tfields []int
		want    []int
	}{
		{"UnionWith test", []int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{"UnionWith test", []int{1, 2}, []int{4, 6}, []int{1, 2, 4, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &IntSet{}
			ts := &IntSet{}
			for _, w := range tt.fields {
				s.Add(w)
			}
			for _, w := range tt.tfields {
				ts.Add(w)
			}
			s.UnionWith(ts)
			if s.Len() != len(tt.want) {
				t.Errorf("got len: %v, want len %v", s.Len(), len(tt.want))
			}
			for _, w := range tt.want {
				if !s.Has(w) {
					t.Errorf("s should have %v", w)
				}
			}
		})
	}
}

func TestIntSet_String(t *testing.T) {
	tests := []struct {
		name   string
		fields []int
		want   string
	}{
		{"String test", []int{1, 2, 3}, "{1 2 3}"},
		{"String test", []int{10, 20, 30}, "{10 20 30}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &IntSet{}
			for _, w := range tt.fields {
				s.Add(w)
			}
			if got := s.String(); got != tt.want {
				t.Errorf("IntSet.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntSet_AddAll(t *testing.T) {
	type args struct {
		num []int
	}
	tests := []struct {
		name   string
		fields []int
		args   args
		want   []int
	}{
		{"AddAll test", []int{1, 2, 3}, args{[]int{4, 5, 6}}, []int{1, 2, 3, 4, 5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &IntSet{}
			for _, w := range tt.fields {
				s.Add(w)
			}
			s.AddAll(tt.args.num...)
			if s.Len() != len(tt.want) {
				t.Errorf("got len: %v, want len %v", s.Len(), len(tt.want))
			}
			for _, w := range tt.want {
				if !s.Has(w) {
					t.Errorf("s should have %v", w)
				}
			}
		})
	}
}

func TestIntSet_IntersectWith(t *testing.T) {
	tests := []struct {
		name   string
		fields []int
		args   []int
		want   []int
	}{
		{"Intersect test", []int{1, 2, 3}, []int{1, 2, 6}, []int{1, 2}},
		{"Intersect test", []int{1, 2, 3, 4, 5, 6}, []int{1, 2, 6}, []int{1, 2, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &IntSet{}
			s.AddAll(tt.fields...)
			ts := &IntSet{}
			ts.AddAll(tt.args...)
			s.IntersectWith(ts)
			if s.Len() != len(tt.want) {
				t.Errorf("got:%v, want:%v", s.String, tt.want)
			}
			for _, w := range tt.want {
				if !s.Has(w) {
					t.Errorf("s doesn't have %v", w)
				}
			}
		})
	}
}

func TestIntSet_DifferenceWith(t *testing.T) {
	tests := []struct {
		name   string
		fields []int
		args   []int
		want   []int
	}{
		{"Intersect test", []int{1, 2, 3}, []int{1, 2}, []int{3}},
		{"Intersect test", []int{1, 2, 3, 4, 5}, []int{1, 2}, []int{3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &IntSet{}
			s.AddAll(tt.fields...)
			ts := &IntSet{}
			ts.AddAll(tt.args...)
			s.DifferenceWith(ts)
			if s.Len() != len(tt.want) {
				t.Errorf("got:%v, want:%v", s.String, tt.want)
			}
			for _, w := range tt.want {
				if !s.Has(w) {
					t.Errorf("s doesn't have %v", w)
				}
			}
		})
	}
}

func TestIntSet_SymmetricDifference(t *testing.T) {
	tests := []struct {
		name   string
		fields []int
		args   []int
		want   []int
	}{
		{"SymmetricDifference test", []int{1, 2, 3}, []int{1, 2}, []int{3}},
		{"SymmetricDifference test", []int{1, 3}, []int{1, 2}, []int{2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &IntSet{}
			s.AddAll(tt.fields...)
			ts := &IntSet{}
			ts.AddAll(tt.args...)
			s.SymmetricDifference(ts)
			if s.Len() != len(tt.want) {
				t.Errorf("got:%v, want:%v", s.String, tt.want)
			}
			for _, w := range tt.want {
				if !s.Has(w) {
					t.Errorf("s doesn't have %v", w)
				}
			}
		})
	}
}
