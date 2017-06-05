package eval

import (
	"fmt"
	"math"
)

type min struct {
	x, y Expr
}

func (m min) String() string {
	return fmt.Sprintf("min(%s, %s)", m.x, m.y)
}

func (m min) Eval(env Env) float64 {
	return math.Min(m.x.Eval(env), m.y.Eval(env))
}

func (m min) Check(vars map[Var]bool) error {
	if err := m.x.Check(vars); err != nil {
		return err
	}
	return m.y.Check(vars)
}
