package vim9gorn

import (
	"fmt"
	"strings"
)

type ClassField struct {
	Name  string
	Type  string
	Value string
}

type Class struct {
	Name    string
	Fields  []ClassField
	Methods []Function
	Supers  []string
}

func NewClass(name string) *Class {
	return &Class{
		Name:    name,
		Fields:  make([]ClassField, 0),
		Methods: make([]Function, 0),
		Supers:  make([]string, 0),
	}
}

func (c *Class) AddField(name, typ string) *Class {
	c.Fields = append(c.Fields, ClassField{
		Name: name,
		Type: typ,
	})
	return c
}

func (c *Class) AddFieldWithDefault(name, typ, value string) *Class {
	c.Fields = append(c.Fields, ClassField{
		Name:  name,
		Type:  typ,
		Value: value,
	})
	return c
}

func (c *Class) AddMethod(fn *Function) *Class {
	c.Methods = append(c.Methods, *fn)
	return c
}

func (c *Class) SetSuper(parent string) *Class {
	c.Supers = append(c.Supers, parent)
	return c
}

func (c *Class) Generate() string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("class %s", c.Name))

	if len(c.Supers) > 0 {
		b.WriteString(" extends ")
		b.WriteString(strings.Join(c.Supers, ", "))
	}
	b.WriteByte('\n')

	for _, field := range c.Fields {
		if field.Value != "" {
			b.WriteString(fmt.Sprintf("    this.%s: %s = %s\n", field.Name, field.Type, field.Value))
		} else {
			b.WriteString(fmt.Sprintf("    this.%s: %s\n", field.Name, field.Type))
		}
	}

	for _, method := range c.Methods {
		methodLines := strings.Split(method.Generate(), "\n")
		for _, line := range methodLines {
			if strings.TrimSpace(line) == "" {
				continue
			}
			b.WriteString("    " + line + "\n")
		}
	}

	b.WriteString("endclass\n")
	return b.String()
}

type ClassInstance struct {
	ClassName string
	Args      []string
}

func NewClassInstance(className string, args ...string) *ClassInstance {
	return &ClassInstance{
		ClassName: className,
		Args:      args,
	}
}

func (ci *ClassInstance) Generate() string {
	if len(ci.Args) == 0 {
		return ci.ClassName + "->new()"
	}
	return fmt.Sprintf("%s->new(%s)", ci.ClassName, strings.Join(ci.Args, ", "))
}
