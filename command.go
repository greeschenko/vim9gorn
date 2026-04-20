package vim9gorn

import (
	"strings"
)

type Command struct {
	Name     string
	Cmd      string
	Range    bool
	Count    bool
	Complete string
	Nargs    string
	Addr     string
	Bang     bool
}

func NewCommand(name, cmd string) *Command {
	return &Command{
		Name: name,
		Cmd:  cmd,
	}
}

func (c *Command) SetRange() *Command {
	c.Range = true
	return c
}

func (c *Command) SetCount() *Command {
	c.Count = true
	return c
}

func (c *Command) SetComplete(complete string) *Command {
	c.Complete = complete
	return c
}

func (c *Command) SetNargs(nargs string) *Command {
	c.Nargs = nargs
	return c
}

func (c *Command) SetBang() *Command {
	c.Bang = true
	return c
}

func (c *Command) Generate() string {
	var b strings.Builder
	b.WriteString("command")

	if c.Bang {
		b.WriteString("!")
	}

	if c.Range {
		b.WriteString(" -range")
	}

	if c.Count {
		b.WriteString(" -count")
	}

	if c.Complete != "" {
		b.WriteString(" -complete=" + c.Complete)
	}

	if c.Nargs != "" {
		b.WriteString(" -nargs=" + c.Nargs)
	}

	b.WriteString(" ")
	b.WriteString(c.Name)
	b.WriteString(" ")
	b.WriteString(c.Cmd)

	return b.String()
}
