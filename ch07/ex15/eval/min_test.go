package eval

import "testing"

func Test_min_String(t *testing.T) {
	type fields struct {
		x Expr
	}
	tests := []struct {
		name   string
		fields fields
		env    Env
		want   float64
	}{
		{
			"test min",
			fields{&min{literal(1.0), literal(2.0)}},
			Env{},
			1.0,
		},
		{
			"test min",
			fields{&min{Var("x"), literal(1.0)}},
			Env{"x": 10.0},
			1.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.x.Eval(tt.env); got != tt.want {
				t.Errorf("got:%v, want:%v", got, tt.want)
			}
		})
	}
}
