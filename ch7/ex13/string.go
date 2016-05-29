// Copyright © 2016 shoarai

// Package eval provides an expression evaluator.
package eval

import "fmt"

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprintf("%f", l)
}

func (u unary) String() string {
	switch u.op {
	case '+', '-':
		return fmt.Sprintf("%c%s", u.op, u.x)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (b binary) String() string {
	switch b.op {
	case '+', '-', '*', '/':
		return fmt.Sprintf("%s %c %s", b.x, b.op, b.y)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

func (c call) String() string {
	switch c.fn {
	case "pow", "sin", "sqrt":
		str := c.fn + "("
		for i, v := range c.args {
			if i != 0 {
				str += ", "
			}
			str += fmt.Sprintf("%s", v)
		}
		return str + ")"
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}
