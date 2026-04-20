package vim9gorn

import (
	"fmt"
	"strings"
)

type Autocmd struct {
	Event   string
	Pattern string
	Cmd     string
	Nested  bool
	Append  bool
}

func NewAutocmd(event, pattern string) *Autocmd {
	return &Autocmd{
		Event:   event,
		Pattern: pattern,
	}
}

func (a *Autocmd) SetCmd(cmd string) *Autocmd {
	a.Cmd = cmd
	return a
}

func (a *Autocmd) SetNested() *Autocmd {
	a.Nested = true
	return a
}

func (a *Autocmd) SetAppend() *Autocmd {
	a.Append = true
	return a
}

func (a *Autocmd) Generate() string {
	var b strings.Builder
	b.WriteString("autocmd")

	if a.Append {
		b.WriteString("!")
	}

	b.WriteString(" ")
	b.WriteString(a.Event)

	if a.Pattern != "" {
		b.WriteString(" ")
		b.WriteString(a.Pattern)
	}

	if a.Nested {
		b.WriteString(" ++nested")
	}

	b.WriteString(" ")
	b.WriteString(a.Cmd)

	return b.String()
}

type AutocmdGroup struct {
	Name   string
	Events []Autocmd
}

func NewAutocmdGroup(name string) *AutocmdGroup {
	return &AutocmdGroup{
		Name:   name,
		Events: make([]Autocmd, 0),
	}
}

func (g *AutocmdGroup) Add(a Autocmd) *AutocmdGroup {
	g.Events = append(g.Events, a)
	return g
}

func (g *AutocmdGroup) Generate() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("augroup %s\n", g.Name))
	b.WriteString("  autocmd!\n")

	for _, a := range g.Events {
		b.WriteString("  ")
		b.WriteString(a.Generate())
		b.WriteByte('\n')
	}

	b.WriteString("augroup END\n")
	return b.String()
}
