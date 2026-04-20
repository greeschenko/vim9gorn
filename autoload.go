package vim9gorn

import (
	"fmt"
	"strings"
)

type AutoloadFunc struct {
	Name       string
	Lib        string
	Args       []FuncArg
	ReturnType string
	Body       []CodeBlock
}

func NewAutoloadFunc(lib, name string) *AutoloadFunc {
	return &AutoloadFunc{
		Lib:  lib,
		Name: name,
		Args: make([]FuncArg, 0),
		Body: make([]CodeBlock, 0),
	}
}

func (f *AutoloadFunc) Arg(name, typ string) *AutoloadFunc {
	f.Args = append(f.Args, FuncArg{Name: name, Type: typ})
	return f
}

func (f *AutoloadFunc) SetReturn(typ string) *AutoloadFunc {
	f.ReturnType = typ
	return f
}

func (f *AutoloadFunc) Add(block CodeBlock) *AutoloadFunc {
	f.Body = append(f.Body, block)
	return f
}

func (f *AutoloadFunc) Generate() string {
	var b strings.Builder

	b.WriteString("def ")
	b.WriteString(f.Lib)
	b.WriteString("#")
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

	b.WriteString("enddef\n")
	return b.String()
}

type AutoloadFile struct {
	Lib   string
	Funcs []AutoloadFunc
}

func NewAutoloadFile(lib string) *AutoloadFile {
	return &AutoloadFile{
		Lib:   lib,
		Funcs: make([]AutoloadFunc, 0),
	}
}

func (af *AutoloadFile) AddFunc(fn AutoloadFunc) *AutoloadFile {
	af.Funcs = append(af.Funcs, fn)
	return af
}

func (af *AutoloadFile) Generate() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("vim9script\n\n# Autoload: %s\n\n", af.Lib))

	for _, fn := range af.Funcs {
		b.WriteString(fn.Generate())
		b.WriteByte('\n')
	}

	return b.String()
}
