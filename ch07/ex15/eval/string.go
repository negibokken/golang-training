package eval

import "fmt"

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprintf("%g", l)
}

func (a unary) String() string {
	return fmt.Sprintf("%c%s", a.op, a.x)
}

func (b binary) String() string {
	return fmt.Sprintf("%s%c%s", b.x, b.op, b.y)
}

func (c call) String() string {
	str := fmt.Sprintf("%s(", c.fn)
	for i, arg := range c.args {
		if i > 0 {
			str += fmt.Sprintf(", ")
		}
		str += arg.String()
	}
	str += fmt.Sprintf(")")
	return str
}
