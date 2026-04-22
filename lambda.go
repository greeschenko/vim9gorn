package vim9gorn

import (
	"fmt"
	"strings"
)

type Lambda struct {
	Args   []string
	Body   string
	Return string
}

func NewLambda(body string) *Lambda {
	return &Lambda{
		Args: make([]string, 0),
		Body: body,
	}
}

func (l *Lambda) Arg(name string) *Lambda {
	l.Args = append(l.Args, name)
	return l
}

func (l *Lambda) SetReturn(typ string) *Lambda {
	l.Return = typ
	return l
}

func (l *Lambda) Generate() string {
	var b strings.Builder

	b.WriteByte('(')
	for i, arg := range l.Args {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(arg)
	}
	b.WriteString(") => ")
	b.WriteString(l.Body)

	if l.Return != "" {
		b.WriteString(": ")
		b.WriteString(l.Return)
	}

	return b.String()
}

func (l *Lambda) GenerateTyped() string {
	var b strings.Builder

	b.WriteByte('(')
	for i, arg := range l.Args {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(arg)
	}
	b.WriteString(") => ")
	b.WriteString(l.Body)

	return b.String()
}

type Closure struct {
	Lambda
	Captures []string
}

func NewClosure(body string) *Closure {
	return &Closure{
		Lambda: Lambda{
			Args: make([]string, 0),
			Body: body,
		},
		Captures: make([]string, 0),
	}
}

func (c *Closure) Arg(name string) *Closure {
	c.Args = append(c.Args, name)
	return c
}

func (c *Closure) SetReturn(typ string) *Closure {
	c.Return = typ
	return c
}

func (c *Closure) AddCapture(name string) *Closure {
	c.Captures = append(c.Captures, name)
	return c
}

func (c *Closure) Generate() string {
	return c.Lambda.Generate()
}

type LambdaCall struct {
	Lambda *Lambda
}

func NewLambdaCall(fn *Lambda) *LambdaCall {
	return &LambdaCall{Lambda: fn}
}

func (lc *LambdaCall) Call(args ...string) string {
	return fmt.Sprintf("%s(%s)", lc.Lambda.Generate(), strings.Join(args, ", "))
}

func ForEach(items string, fn *Lambda) *ForLoop {
	return NewForLoop("_", "item", items).
		Add(Raw{Code: fn.Generate() + "(item)"})
}

func Filter(items string, pred *Lambda) string {
	return fmt.Sprintf("filter(%s, %s)", items, pred.Generate())
}

func Map(items string, fn *Lambda) string {
	return fmt.Sprintf("map(%s, %s)", items, fn.Generate())
}
