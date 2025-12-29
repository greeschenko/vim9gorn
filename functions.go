package vim9gorn

import (
	"strings"
)

/*
CodeBlock represents any element
that can generate Vim9 code.
Indentation is handled by the parent (Function).
*/
type CodeBlock interface {
	Generate() string
}

/*
FuncArg represents a Vim9 function argument.
Example:

	name: string
*/
type FuncArg struct {
	Name string
	Type string
}

/*
Function represents a Vim9 function (def).
*/
type Function struct {
	Name       string
	Args       []FuncArg
	ReturnType string
	Body       []CodeBlock
	Scope      Scope
}

/*
NewFunction creates a new Vim9 function.
*/
func NewFunction(name string) *Function {
	return &Function{
		Name:  name,
		Args:  make([]FuncArg, 0),
		Body:  make([]CodeBlock, 0),
		Scope: None,
	}
}

func (f *Function) SetScope(s Scope) *Function {
	f.Scope = s
	return f
}

/*
Arg adds a function argument.
*/
func (f *Function) Arg(name, typ string) *Function {
	f.Args = append(f.Args, FuncArg{
		Name: name,
		Type: typ,
	})
	return f
}

/*
Returns sets the function return type.
*/
func (f *Function) Returns(typ string) *Function {
	f.ReturnType = typ
	return f
}

/*
Add appends a CodeBlock to the function body.
*/
func (f *Function) Add(block CodeBlock) *Function {
	f.Body = append(f.Body, block)
	return f
}

/*
Generate generates a Vim9 def function.
Indentation is applied automatically to the body.
*/
func (f *Function) Generate() string {
	var b strings.Builder

	// Function signature
	b.WriteString("def ")

	if f.Scope != None {
		b.WriteString(string(f.Scope))
		b.WriteString(":")
	}

	b.WriteString(f.Name)
	b.WriteString("(")

	for i, arg := range f.Args {
		if i > 0 {
			b.WriteString(", ")
		}
		if arg.Type != "" {
			b.WriteString(arg.Name + ": " + arg.Type)
		} else {
			b.WriteString(arg.Name)
		}
	}

	b.WriteString(")")
	if f.ReturnType != "" {
		b.WriteString(": " + f.ReturnType)
	}
	b.WriteByte('\n')

	// Function body
	for _, block := range f.Body {
		code := block.Generate()
		lines := strings.Split(code, "\n")

		for _, line := range lines {
			if strings.TrimSpace(line) == "" {
				continue
			}
			b.WriteString(functionIndent)
			b.WriteString(line)
			b.WriteByte('\n')
		}
	}

	// End function
	b.WriteString("enddef\n\n")
	return b.String()
}
