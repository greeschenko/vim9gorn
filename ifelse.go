package vim9gorn

import (
	"fmt"
	"strings"
)

/*
IfElse represents an if/elseif/else block in Vim9.
*/
type IfElse struct {
	Condition string
	Then      []CodeBlock
	ElseIfs   []ElseIf
	Else      []CodeBlock
}

/*
ElseIf represents an elseif branch.
*/
type ElseIf struct {
	Condition string
	Body      []CodeBlock
}

/*
NewIfElse creates a new IfElse block with a main condition.
*/
func NewIfElse(cond string) *IfElse {
	return &IfElse{
		Condition: cond,
		Then:      make([]CodeBlock, 0),
		ElseIfs:   make([]ElseIf, 0),
		Else:      make([]CodeBlock, 0),
	}
}

/*
ThenAdd appends a CodeBlock to the main 'then' branch.
*/
func (i *IfElse) ThenAdd(block CodeBlock) *IfElse {
	i.Then = append(i.Then, block)
	return i
}

/*
ElseIfAdd appends an elseif branch.
*/
func (i *IfElse) ElseIfAdd(cond string, body ...CodeBlock) *IfElse {
	i.ElseIfs = append(i.ElseIfs, ElseIf{
		Condition: cond,
		Body:      body,
	})
	return i
}

/*
ElseAdd appends CodeBlocks to the else branch.
*/
func (i *IfElse) ElseAdd(blocks ...CodeBlock) *IfElse {
	i.Else = append(i.Else, blocks...)
	return i
}

/*
Generate produces Vim9 if/elseif/else code.
*/
func (i *IfElse) Generate() string {
	var b strings.Builder

	// main if
	b.WriteString(fmt.Sprintf("if %s\n", i.Condition))
	for _, block := range i.Then {
		lines := strings.Split(block.Generate(), "\n")
		for _, line := range lines {
			if strings.TrimSpace(line) == "" {
				continue
			}
			b.WriteString("  " + line + "\n")
		}
	}

	// elseif branches
	for _, ei := range i.ElseIfs {
		b.WriteString(fmt.Sprintf("elseif %s\n", ei.Condition))
		for _, block := range ei.Body {
			lines := strings.Split(block.Generate(), "\n")
			for _, line := range lines {
				if strings.TrimSpace(line) == "" {
					continue
				}
				b.WriteString("  " + line + "\n")
			}
		}
	}

	// else branch
	if len(i.Else) > 0 {
		b.WriteString("else\n")
		for _, block := range i.Else {
			lines := strings.Split(block.Generate(), "\n")
			for _, line := range lines {
				if strings.TrimSpace(line) == "" {
					continue
				}
				b.WriteString("  " + line + "\n")
			}
		}
	}

	b.WriteString("endif\n")
	return b.String()
}
